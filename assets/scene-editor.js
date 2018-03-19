$(document).ready(function () {
    const baseUrl = window.location.protocol + "//" + window.location.host
    const container = $('#scene-editor .timeline').get(0)
    const selectDropdown = $('#deviceSelection')

    let idx = document.location.href.lastIndexOf("/")

    document.timeline = null
    let devices = null

    function formatTime(d) {
        return d.toLocaleTimeString('de-DE', {
            hour: '2-digit',
            minute: '2-digit',
            // second: '2-digit'
        })
    }

    function epochTime(sTime) {
        return Date.parse('1970-01-01T' + sTime)
    }

    function toVisjsEvent(event) {
        // convert list of property maps into an Object
        if (event.props != null) {
            let propsObject = new Object()
            event.props.forEach((p) => {
                propsObject[p.name] = p.value
            })
            event.props = propsObject
        }
        return event
    }

    function epochToTime(seconds) {
        let d = new Date(0)
        d.setUTCSeconds(seconds)
        return d
    }

    function toServerEvent(event) {
        if (typeof event.start === 'object') {
            event.start = event.start.getTime()
        }
        if (typeof event.end === 'object') {
            event.end = event.end.getTime()
        }
        if (event.props != null) {
            // convert property Object into a list of property maps
            let propsArray = []
            Object.keys(event.props).forEach((k) => {
                propsArray.push(new Object({ "name": k, "value": event.props[k] }))
            })
            event.props = propsArray
        }
        return event
    }

    function zipmap(devices, events) {
        let deviceLookup = new Map()
        let deviceEvents = new Map()
        devices.forEach((d) => {
            deviceLookup.set(d.id, d)
            deviceEvents.set(d, new Array())
        })
        events.forEach((e) => {
            let device = deviceLookup.get(e.group)
            deviceEvents.get(device).push(e)
        })
        return deviceEvents
    }

    function split(deviceEvents) {
        let activeDeviceEvents = new Map()
        let inactiveDevices = new Array()
        deviceEvents.forEach((value, key, map) => {
            if (value.length == 0) {
                inactiveDevices.push(key)
            } else {
                activeDeviceEvents.set(key, value)
            }
        })
        return {
            activeDeviceEvents: activeDeviceEvents,
            inactiveDevices: inactiveDevices
        }
    }

    function updateDeviceSelection(selectElement, timeline) {
        if (selectElement != null && devices != null) {
            let activeDeviceIDs = timeline.groupsData.map(x => x.id)
            selectElement.empty()
            selectElement.append($("<option></option>").attr("value", "").attr("disabled", true).attr("selected", true).text("Hinzufügen..."));
            devices.forEach((d) => {
                if (activeDeviceIDs.indexOf(d.id) == -1) {
                    selectElement.append($("<option></option>").attr("value", d.id).text(d.content));
                }
            })
        }
    }

    $('#deviceSelection').change(
        function (e) {
            let target = e.target
            let deviceID = target[target.selectedIndex].value
            console.log(`Add device ID ${deviceID}`)
            addDevice(deviceID)
        }
    )

    $(document).on("click", ".remove_device",
        function (e) {
            let deviceID = e.target.getAttribute('group')
            console.log(`Remove device ID ${deviceID}`)
            removeDevice(deviceID)
        }
    )

    $(document).on("click", "#editorSubmit",
        function (e) {
            console.log("Save Event Action!")

            data = {
                "scene": sceneID,
                "group": $('#duft-edit input[name="group"]').val(),
                "start": epochTime($('#duft-edit input[name="start"]').val()),
                "end": epochTime($('#duft-edit input[name="end"]').val()),
                "props": {
                    "title": $('#duft-edit input[name="title"]').val(),
                    "description": $('#duft-edit input[name="description"]').val(),
                    "color": $('#duft-edit input[name="color"]:checked').val(),
                }
            }

            // TODO: Get callback function stored before opening modal window
            callback = window.cda_visjs_callback

            let eventID = $('#duft-edit input[name="id"]').val()
            if (eventID === "") {
                addEvent(data, callback)
            } else {
                data.id = eventID
                updateEvent(data, callback)
            }

            $('#duft-edit').modal('hide')
        }
    )

    function showActionEditor(item, isNew, callback) {
        $('#duft-edit').modal('show')
        // $('#duft-edit').on('hidden.bs.modal', function (e) {
        //     console.log("modal hidden.bs.modal event:")
        //     //e.preventDefault()
        //     console.log(e.data)
        //     alert(JSON.stringify(e.data))
        //     return true
        // })
        if (isNew) {
            item.end = Date.parse(item.start) + 5400000
            item.props = {
                title: "",
                description: ""
            }
            $('#duft-edit input[name="id"]').val("")
        } else {
            $('#duft-edit input[name="id"]').val(item.id)
        }

        // TODO: Store callback fn for conversation...
        window.cda_visjs_callback = callback

        $('#duft-edit input[name="group"]').val(item.group)
        $('#duft-edit input[name="title"]').val(item.props.title)
        $('#duft-edit input[name="description"]').val(item.props.description)
        $('#duft-edit input[name="start"]').val(formatTime(new Date(item.start)))
        $('#duft-edit input[name="end"]').val(formatTime(new Date(item.end)))
        if (item.props.color !== undefined) {
            $(`#duft-edit input[name="color"][value="${item.props.color}"]`).prop("checked", true)
        }
    }

    // ------------------------------------------------------------------------------------

    function addDevice(deviceID) {
        console.log(`Add device ID ${deviceID}`)
        let device = devices.find(d => { return (d.id === deviceID) })
        if (device != null) {
            $.ajax({
                url: `${baseUrl}/scene/data/${sceneID}/device/${deviceID}`,
                context: this,
                method: 'PUT',
                contentType: 'application/json',
                data: JSON.stringify(device),
                success: function (data) {
                    timeline.groupsData.getDataSet().add(device)
                    updateDeviceSelection(selectDropdown, timeline)
                },
                error: function (err) {
                    console.log('Error', err);
                    if (err.status === 0) {
                        alert('Failed to load actions.\nPlease run this example on a server.');
                    }
                    else {
                        alert('Failed to load actions.');
                    }
                }
            })
        }
    }

    function removeDevice(deviceID) {
        console.log(`Remove device ID ${deviceID}`)

        let devices = timeline.groupsData.getDataSet().get({
            filter: function (d) {
                return (d.id === deviceID)
            }
        })

        if (devices != null && devices.length == 1) {
            $.ajax({
                url: `${baseUrl}/scene/data/${sceneID}/device/${deviceID}`,
                context: this,
                method: 'DELETE',
                contentType: 'application/json',
                success: function (data) {
                    timeline.groupsData.getDataSet().remove(devices[0])
                    updateDeviceSelection(selectDropdown, timeline)
                },
                error: function (err) {
                    console.log('Error', err);
                    if (err.status === 0) {
                        alert('Failed to load actions.\nPlease run this example on a server.');
                    }
                    else {
                        alert('Failed to load actions.');
                    }
                }
            })
        }
    }

    function addEvent(event, callback) {
        console.log("SceneEditor add event: " + JSON.stringify(event))
        event.scene_id = sceneID

        let srvEvent = toServerEvent(event)
        $.ajax({
            url: `${baseUrl}/scene/data/${sceneID}/event`,
            context: this,
            method: 'POST',
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            data: JSON.stringify(srvEvent),
            processData: false,
            success: function (data) {
                let visEvent = toVisjsEvent(data)
                console.log("Add event: " + JSON.stringify(visEvent))
                callback(visEvent)
            },
            error: function (err) {
                console.log('Error', err);
                if (err.status === 0) {
                    alert('Failed to load actions.\nPlease run this example on a server.');
                }
                else {
                    alert('Failed to load actions.');
                }
            }
        })
    }

    function updateEvent(event, callback) {
        console.log("SceneEditor update event: " + JSON.stringify(event))

        let events = timeline.itemsData.getDataSet().get({
            filter: function (e) {
                return (e.id === event.id)
            }
        })

        let srvEvent = toServerEvent(event)
        $.ajax({
            url: `${baseUrl}/scene/data/${sceneID}/event/${event.id}`,
            context: this,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(srvEvent),
            success: function (data) {
                let visEvent = toVisjsEvent(data)
                console.log("Updated event: " + JSON.stringify(visEvent))
                callback(visEvent)
            },
            error: function (err) {
                console.log('Error', err);
                if (err.status === 0) {
                    alert('Failed to load actions.\nPlease run this example on a server.');
                }
                else {
                    alert('Failed to load actions.');
                }
            }
        })
    }


    function removeEvent(event, callback) {
        let eventID = event.id
        console.log(`Remove event ID ${eventID}`)

        let events = timeline.itemsData.getDataSet().get({
            filter: function (e) {
                return (e.id === eventID)
            }
        })

        if (events != null && events.length == 1) {
            $.ajax({
                url: `${baseUrl}/scene/data/${sceneID}/event/${eventID}`,
                context: this,
                method: 'DELETE',
                contentType: 'application/json',
                success: function (data) {
                    let visEvent = toVisjsEvent(data)
                    console.log("Remove event: " + JSON.stringify(visEvent))
                    callback(visEvent)
                },
                error: function (err) {
                    console.log('Error', err);
                    if (err.status === 0) {
                        alert('Failed to load actions.\nPlease run this example on a server.');
                    }
                    else {
                        alert('Failed to load actions.');
                    }
                }
            })
        }
    }

    // ------------------------------------------------------------------------------------

    const timelineOptions = {
        min: 0,
        max: 86400000,
        start: 0,
        end: 86400000,

        editable: {
            add: true,            // add new items by double tapping
            updateTime: true,     // drag items horizontally
            updateGroup: false,   // drag items from one group to another
            remove: true,         // delete an item by tapping the delete button top right
            overrideItems: false  // allow these options to override item.editable
        },

        showMajorLabels: false,

        template: function (item, element, data) {
            return `<b> ${item.props.title}</b><br/> ${item.props.description}`;
        },

        groupOrder: "content",

        tooltipOnItemUpdateTime: {
            template: function (item) {
                return item.start.toLocaleTimeString('de-CH', { hour: "numeric", minute: "numeric" }) + ' -> ' + item.end.toLocaleTimeString('de-CH', { hour: "numeric", minute: "numeric" });
            }
        },

        showTooltips: false,

        snap: null,

        // snap: function (date, scale, step) {
        //     var hour = 60 * 60 * 1000;
        //     return Math.round(date / hour) * hour;
        // },

        // onMoving: function (item, callback) {
        //     item.moving = true;
        // },

        onUpdate: function (item, callback) {
            showActionEditor(item, false, callback)
        },

        onAdd: function (item, callback) {
            showActionEditor(item, true, callback)
        },

        onRemove: function (item, callback) {
            removeEvent(item, callback)
        },

        onMove: function (item, callback) {
            updateEvent(item, callback)
        },

        groupTemplate: function (group) {
            let container = document.createElement('div');
            let label = document.createElement('div');
            label.innerHTML = group.content + ' ';
            container.insertAdjacentElement('afterBegin', label);
            let hide = document.createElement('button')
            hide.setAttribute('group', group.id)
            hide.classList.add('btn', 'btn-default', 'btn-xs', 'remove_device');
            hide.innerHTML = 'Remove';
            hide.style.fontSize = '0.7em';
            container.insertAdjacentElement('beforeEnd', hide);

            return container;
        },
    }

    // Initialization
    console.log(`Load scene data for ID ${sceneID} (baseUrl: ${baseUrl})`)
    $.ajax({
        url: `${baseUrl}/scene/data/${sceneID}`,
        context: this,
        success: function (data) {
            devices = data.devices
            let events = data.events.map(toVisjsEvent)
            let deviceEvents = zipmap(devices, events)
            let { activeDeviceEvents, inactiveDevices } = split(deviceEvents)

            timeline = new vis.Timeline(
                container,
                new vis.DataSet(Array.from(activeDeviceEvents.values()).reduce((acc, events) => acc.concat(events))),
                new vis.DataSet(Array.from(new Map(Array.from(deviceEvents.entries()).filter(([k, v]) => v.length > 0)).keys())),
                timelineOptions
            )

            updateDeviceSelection(selectDropdown, timeline)
            console.log(`Successfully initialized Timeline`)
        },
        error: function (err) {
            console.log('Error', err);
            if (err.status === 0) {
                alert('Failed to load actions.\nPlease run this example on a server.');
            }
            else {
                alert('Failed to load actions.');
            }
        }
    })

});



