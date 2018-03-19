package grifts

import (
	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/pop"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		return nil
	})

	grift.Desc("seed:scene", "Seed scene data")
	grift.Add("seed:scene", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			if err := tx.TruncateAll(); err != nil {
				return errors.WithStack(err)
			}

			// Scene: scene

			scene := &models.Scene{
				Name: "Scene Example",
			}
			if err := tx.Create(scene); err != nil {
				return errors.WithStack(err)
			}

			// Devices: light1, light2, video, scent1, scent2

			// light1
			light1 := &models.Device{
				Name: "Osram Leuchtband",
				Type: "Light",
			}
			if err := tx.Create(light1); err != nil {
				return errors.WithStack(err)
			}

			// light2
			light2 := &models.Device{
				Name: "Philips E27",
				Type: "Light",
			}
			if err := tx.Create(light2); err != nil {
				return errors.WithStack(err)
			}

			// light3
			light3 := &models.Device{
				Name: "LED Oberlicht",
				Type: "Light",
			}
			if err := tx.Create(light3); err != nil {
				return errors.WithStack(err)
			}

			// video
			video := &models.Device{
				Name: "Video",
				Type: "Video",
			}
			if err := tx.Create(video); err != nil {
				return errors.WithStack(err)
			}

			// scent1
			scent1 := &models.Device{
				Name: "Duft Links",
				Type: "Scent",
			}
			if err := tx.Create(scent1); err != nil {
				return errors.WithStack(err)
			}

			// scent2
			scent2 := &models.Device{
				Name: "Duft Oben",
				Type: "Scent",
			}
			if err := tx.Create(scent2); err != nil {
				return errors.WithStack(err)
			}

			// scent3
			scent3 := &models.Device{
				Name: "Duft Hinten",
				Type: "Scent",
			}
			if err := tx.Create(scent3); err != nil {
				return errors.WithStack(err)
			}

			// Events

			// light1
			//   event1

			e1 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: light1.ID,
				Start:    14400000,
				End:      18960000,
			}
			if err := tx.Create(e1); err != nil {
				return errors.WithStack(err)
			}

			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e1.ID,
					Name:    "title",
					Value:   "Frühlingserwachen",
				},
				models.Prop{
					EventID: e1.ID,
					Name:    "description",
					Value:   "Montag",
				},
				models.Prop{
					EventID: e1.ID,
					Name:    "color",
					Value:   "red",
				},
				models.Prop{
					EventID: e1.ID,
					Name:    "saturation",
					Value:   "159",
				},
				models.Prop{
					EventID: e1.ID,
					Name:    "brightness",
					Value:   "75",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// light1
			//   event2
			e2 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: light1.ID,
				Start:    33300000,
				End:      40560000,
			}
			if err := tx.Create(e2); err != nil {
				return errors.WithStack(err)
			}
			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e2.ID,
					Name:    "title",
					Value:   "Tanz",
				},
				models.Prop{
					EventID: e2.ID,
					Name:    "description",
					Value:   "Himmelspiel",
				},
				models.Prop{
					EventID: e2.ID,
					Name:    "color",
					Value:   "blue",
				},
				models.Prop{
					EventID: e2.ID,
					Name:    "saturation",
					Value:   "200",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// light2
			//   event3
			e3 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: light2.ID,
				Start:    61200000,
				End:      77400000,
			}
			if err := tx.Create(e3); err != nil {
				return errors.WithStack(err)
			}

			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e3.ID,
					Name:    "title",
					Value:   "Wald",
				},
				models.Prop{
					EventID: e3.ID,
					Name:    "description",
					Value:   "Urwald",
				},
				models.Prop{
					EventID: e3.ID,
					Name:    "color",
					Value:   "green",
				},
				models.Prop{
					EventID: e3.ID,
					Name:    "brightness",
					Value:   "24",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// video
			//   event4

			e4 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: video.ID,
				Start:    57600000,
				End:      69360000,
			}
			if err := tx.Create(e4); err != nil {
				return errors.WithStack(err)
			}
			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e4.ID,
					Name:    "title",
					Value:   "Meeresbucht",
				},
				models.Prop{
					EventID: e4.ID,
					Name:    "description",
					Value:   "Wellenspiel",
				},
				models.Prop{
					EventID: e4.ID,
					Name:    "media",
					Value:   "61c533dc-a540-b5d4-904a-31a9a0d567aa",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// scent1
			//   event5
			e5 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent1.ID,
				Start:    35400000,
				End:      48900000,
			}
			if err := tx.Create(e5); err != nil {
				return errors.WithStack(err)
			}
			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e5.ID,
					Name:    "title",
					Value:   "Garten",
				},
				models.Prop{
					EventID: e5.ID,
					Name:    "description",
					Value:   "Rosenduft",
				},
				models.Prop{
					EventID: e5.ID,
					Name:    "scent",
					Value:   "blue",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// scent1
			//    event6
			e6 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent1.ID,
				Start:    50400000,
				End:      57600000,
			}
			if err := tx.Create(e6); err != nil {
				return errors.WithStack(err)
			}
			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e6.ID,
					Name:    "title",
					Value:   "Space",
				},
				models.Prop{
					EventID: e6.ID,
					Name:    "description",
					Value:   "Weltall",
				},
				models.Prop{
					EventID: e6.ID,
					Name:    "delay",
					Value:   "1200",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			// scent2
			//   event7
			e7 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent2.ID,
				Start:    72400000,
				End:      81300000,
			}
			if err := tx.Create(e7); err != nil {
				return errors.WithStack(err)
			}
			for _, prop := range []models.Prop{
				models.Prop{
					EventID: e7.ID,
					Name:    "title",
					Value:   "See",
				},
				models.Prop{
					EventID: e7.ID,
					Name:    "description",
					Value:   "Algen",
				},
				models.Prop{
					EventID: e7.ID,
					Name:    "scent",
					Value:   "magenta",
				}} {
				if err := tx.Create(&prop); err != nil {
					return errors.WithStack(err)
				}
			}

			return nil
		})
	})
})
