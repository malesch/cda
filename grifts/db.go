package grifts

import (
	"github.com/cdacontrol/cda/models"
	"github.com/markbates/grift/grift"
	"github.com/markbates/pop"
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
			err := tx.TruncateAll()
			if err != nil {
				return errors.WithStack(err)
			}

			scene := &models.Scene{
				Name: "Scene Example",
			}
			err = tx.Create(scene)
			if err != nil {
				return errors.WithStack(err)
			}

			led := &models.Device{
				Name: "LED",
				Type: "LED Device",
			}
			err = tx.Create(led)
			if err != nil {
				return errors.WithStack(err)
			}

			video := &models.Device{
				Name: "Video",
				Type: "Video",
			}
			err = tx.Create(video)
			if err != nil {
				return errors.WithStack(err)
			}

			scent := &models.Device{
				Name: "Duft",
				Type: "Scent Device",
			}
			err = tx.Create(scent)
			if err != nil {
				return errors.WithStack(err)
			}

			e1 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: led.ID,
				Start:    14400000,
				End:      18960000,
			}
			err = tx.Create(e1)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e1.ID,
				Name:    "color",
				Value:   "red",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e2 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: led.ID,
				Start:    33300000,
				End:      40560000,
			}
			err = tx.Create(e2)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e2.ID,
				Name:    "color",
				Value:   "blue",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e2.ID,
				Name:    "fadeIn",
				Value:   "2000",
				Type:    "long",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e3 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: led.ID,
				Start:    61200000,
				End:      77400000,
			}
			err = tx.Create(e3)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e3.ID,
				Name:    "color",
				Value:   "green",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e4 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: video.ID,
				Start:    57600000,
				End:      69360000,
			}
			err = tx.Create(e4)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e4.ID,
				Name:    "mediaID",
				Value:   "61c533dc-a540-b5d4-904a-31a9a0d567aa",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e5 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent.ID,
				Start:    35400000,
				End:      48900000,
			}
			err = tx.Create(e5)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e5.ID,
				Name:    "scent",
				Value:   "blue",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e6 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent.ID,
				Start:    50400000,
				End:      57600000,
			}
			err = tx.Create(e6)
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.Create(&models.Prop{
				EventID: e6.ID,
				Name:    "scent",
				Value:   "orange",
				Type:    "string",
			})
			if err != nil {
				return errors.WithStack(err)
			}

			e7 := &models.Event{
				SceneID:  scene.ID,
				DeviceID: scent.ID,
				Start:    72400000,
				End:      81300000,
			}
			err = tx.Create(e7)
			if err != nil {
				return errors.WithStack(err)
			}

			return tx.Create(&models.Prop{
				EventID: e7.ID,
				Name:    "scent",
				Value:   "magenta",
				Type:    "string",
			})
		})
	})
})
