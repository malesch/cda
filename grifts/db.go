package grifts

import (
	"github.com/cdacontrol/cda/models"
	"github.com/markbates/grift/grift"
	"github.com/markbates/pop/nulls"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		return nil
	})

	grift.Desc("scene", "Seed scene data")
	grift.Add("scene", func(c *grift.Context) error {
		sceneIconMedium := &models.Medium{
			Name: nulls.NewString("Scene Example Icon"),
			Type: "image/png",
			Size: 123123,
		}
		models.DB.Create(sceneIconMedium)

		scene := &models.Scene{
			Name:     "Scene Example",
			MediumID: sceneIconMedium.ID,
		}
		models.DB.Create(scene)

		led := &models.Device{
			Name: "LED",
			Type: "LED Device",
		}
		models.DB.Create(led)

		video := &models.Device{
			Name: "Video",
			Type: "Video",
		}
		models.DB.Create(video)

		scent := &models.Device{
			Name: "Duft",
			Type: "Scent Device",
		}
		models.DB.Create(scent)

		e1 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: led.ID,
			Start:    14400000,
			End:      18960000,
		}
		models.DB.Create(e1)
		models.DB.Create(&models.Prop{
			EventID: e1.ID,
			Name:    "color",
			Value:   "red",
			Type:    "string",
		})

		e2 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: led.ID,
			Start:    33300000,
			End:      40560000,
		}
		models.DB.Create(e2)
		models.DB.Create(&models.Prop{
			EventID: e2.ID,
			Name:    "color",
			Value:   "blue",
			Type:    "string",
		})
		models.DB.Create(&models.Prop{
			EventID: e2.ID,
			Name:    "fadeIn",
			Value:   "2000",
			Type:    "long",
		})

		e3 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: led.ID,
			Start:    61200000,
			End:      77400000,
		}
		models.DB.Create(e3)
		models.DB.Create(&models.Prop{
			EventID: e3.ID,
			Name:    "color",
			Value:   "green",
			Type:    "string",
		})

		e4 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: video.ID,
			Start:    57600000,
			End:      69360000,
		}
		models.DB.Create(e4)
		models.DB.Create(&models.Prop{
			EventID: e4.ID,
			Name:    "mediaID",
			Value:   "61c533dc-a540-b5d4-904a-31a9a0d567aa",
			Type:    "string",
		})

		e5 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: scent.ID,
			Start:    35400000,
			End:      48900000,
		}
		models.DB.Create(e5)
		models.DB.Create(&models.Prop{
			EventID: e5.ID,
			Name:    "scent",
			Value:   "blue",
			Type:    "string",
		})

		e6 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: scent.ID,
			Start:    50400000,
			End:      57600000,
		}
		models.DB.Create(e6)
		models.DB.Create(&models.Prop{
			EventID: e6.ID,
			Name:    "scent",
			Value:   "orange",
			Type:    "string",
		})

		e7 := &models.Event{
			SceneID:  scene.ID,
			DeviceID: scent.ID,
			Start:    72400000,
			End:      81300000,
		}
		models.DB.Create(e7)
		models.DB.Create(&models.Prop{
			EventID: e7.ID,
			Name:    "scent",
			Value:   "magenta",
			Type:    "string",
		})

		return nil
	})
})
