// DOM element where the Timeline will be attached
// var container = $('#scene-editor .timeline').get(0);

function formatTime(d) {
    return d.toLocaleTimeString('de-DE', {
        hour: '2-digit',
        minute: '2-digit',
        //    second: '2-digit'
    })
}

function epochTime(sTime) {
    return Date.parse('1970-01-01T' + sTime)
}

class SceneEditor {

    constructor(container, scene_id) {
        this.scene_id = scene_id;
        this.baseUrl = window.location.protocol + "//" + window.location.host 
        this.timeline = new vis.Timeline(container);
        this.timeline.setOptions({
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
                //console.log(item);
                return '<b>' + item.title + '</b><br/>' + item.description;
            },

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
                this.showActionEditor(item, false, callback)
            },

            onAdd: function (item, callback) {
                console.log(item)
                this.showActionEditor(item, true, callback)
            },

            onRemove: function (item, callback) {
                //alert("Delete action ID " + item.id)
                $.ajax({
                    url: `{this.baseUrl}/scene/` + item.id,
                    method: 'DELETE',
                    success: function (data) {
                        callback(data)
                    },
                    error: function (err) {
                        console.log('Error', err);
                        if (err.status === 0) {
                            alert('Failed to load actions.\nPlease run this example on a server.');
                        }
                        else {
                            alert('Failed to load actions.');
                        }
                        callback(null)
                    }

                })
            }
        });

        this.data_options = {};
        this.groupData = new vis.DataSet(this.data_options)
        this.groupDataCB = this.groupData.on('*', function (event, properties, senderId) {
            console.log('group update: ' + JSON.stringify(properties.data))
        });
        this.itemData = new vis.DataSet(this.data_options)
        this.groupDataDB = this.itemData.on('*', function (event, properties, senderId) {
            console.log('action update: ' + JSON.stringify(properties.data))
            if (event == 'update') {
                data = properties.data[0]
                data.start = Date.parse(data.start)
                data.end = Date.parse(data.end)
                const newLocal = JSON.stringify(data);

                $.ajax({
                    url: `{this.baseUrl}/scene/${data.id}`,
                    method: 'PUT',
                    contentType: 'application/json',
                    data: newLocal,
                    success: function (data) {
                        console.log("added: " + data)
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
        });
    }

    init() {
        $.ajax({
            url: `http://localhost:3000/scene/data/${this.scene_id}`,
            success: function (data) {
                console.log("hoho devices = " + data.devices)
                console.log("hoho events  = " + data.events)
                
                this.groupData.add(data.devices)
                this.itemData.add(data.events)

                this.timeline.setGroups(this.groupData)
                this.timeline.setItems(this.itemData)
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
        });
    }

    showActionEditor(item, isNew, callback) {
        if (isNew) {
            item.end = Date.parse(item.start) + 5400000
            // item.title = 'New entry'
            // item.description = 'nix'
        } else {
            $('#duft-edit input[name="id"]').val(item.id)
        }
        console.log(item)
        $('#duft-edit input[name="group"]').val(item.group)
        $('#duft-edit input[name="title"]').val(item.title)
        $('#duft-edit input[name="description"]').val(item.description)
        $('#duft-edit input[name="start"]').val(formatTime(new Date(item.start)))
        $('#duft-edit input[name="end"]').val(formatTime(new Date(item.end)))
        $('#duft-edit').modal('show')

        callback(null)
    }

    saveUpdateAction(item) {
        console.log("Save Action!")

        id = $('#duft-edit input[name="id"]').val()
        data = {
            "group": $('#duft-edit input[name="group"]').val(),
            "title": $('#duft-edit input[name="title"]').val(),
            "description": $('#duft-edit input[name="description"]').val(),
            "start": epochTime($('#duft-edit input[name="start"]').val()),
            "end": epochTime($('#duft-edit input[name="end"]').val()),
        }

        if (id == "") {
            console.log("Create action")
            $.ajax({
                url: `{this.baseUrl}/scene`,
                method: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(data),
                success: function (data) {
                    console.log("added: " + data)
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
        } else {
            console.log("Update action")
            $.ajax({
                url: `{this.baseUrl}/scene/${this.scene_id}`,
                method: 'PUT',
                contentType: 'application/json',
                data: JSON.stringify(data),
                success: function (data) {
                    console.log("added: " + data)
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
        $('#duft-edit').hide('show')
    }
}