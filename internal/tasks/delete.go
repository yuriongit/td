/*
Package tasks offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package tasks

import (
	"context"

	"github.com/yuriongit/td/internal/app"
	"github.com/yuriongit/td/internal/domain"
	"github.com/yuriongit/td/internal/ui"
)

// Remove removes a task by ID.
func Remove(rootCtx context.Context, id domain.ID, db *app.SupabaseDB) error {
	if err := validateID(id); err != nil {
		return err
	}

	if err := db.DeleteTask(rootCtx, id); err != nil {
		return err
	}

	ui.ColoredTaskOutput.Remove(id)

	return nil
}
