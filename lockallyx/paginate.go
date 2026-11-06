package lockallyx

import "context"

// Page is one page of a cursor-paginated list. NextCursor is empty on the last
// page. The Lockally API returns `next_cursor`; adapt the generated list
// response into this shape inside the fetch func.
type Page[T any] struct {
	Items      []T
	NextCursor string
}

// Paginate walks every page of a cursor-paginated endpoint, calling yield for
// each item. Return false from yield to stop early. fetch is called with the
// cursor for the next page ("" for the first).
//
//	err := lockallyx.Paginate(ctx,
//	  func(ctx context.Context, cursor string) (lockallyx.Page[lockally.Message], error) {
//	    res, _, err := client.MessagesAPI.ListMessages(ctx).Cursor(cursor).Execute()
//	    if err != nil { return lockallyx.Page[lockally.Message]{}, err }
//	    return lockallyx.Page[lockally.Message]{Items: res.Data, NextCursor: res.GetNextCursor()}, nil
//	  },
//	  func(m lockally.Message) bool { fmt.Println(m.GetId()); return true },
//	)
func Paginate[T any](
	ctx context.Context,
	fetch func(ctx context.Context, cursor string) (Page[T], error),
	yield func(T) bool,
) error {
	cursor := ""
	seen := map[string]bool{}
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		page, err := fetch(ctx, cursor)
		if err != nil {
			return err
		}
		for _, item := range page.Items {
			if !yield(item) {
				return nil
			}
		}
		if page.NextCursor == "" || page.NextCursor == cursor || seen[page.NextCursor] {
			return nil // done, or a misbehaving server looping the cursor
		}
		seen[cursor] = true
		cursor = page.NextCursor
	}
}

// Collect drains every page into a single slice. Convenient for small result
// sets; prefer Paginate for large ones to avoid holding everything in memory.
func Collect[T any](
	ctx context.Context,
	fetch func(ctx context.Context, cursor string) (Page[T], error),
) ([]T, error) {
	var out []T
	err := Paginate(ctx, fetch, func(item T) bool {
		out = append(out, item)
		return true
	})
	return out, err
}
