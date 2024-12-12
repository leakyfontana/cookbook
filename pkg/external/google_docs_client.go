package external

import (
	"context"
	"fmt"

	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func CreateGoogleDoc(title string, content string) (string, error) {
    ctx := context.Background()
    srv, err := docs.NewService(ctx, option.WithAPIKey("YOUR_GOOGLE_API_KEY"))
    if err != nil {
        return "", err
    }

    doc := &docs.Document{Title: title}
    createdDoc, err := srv.Documents.Create(doc).Do()
    if err != nil {
        return "", err
    }

    // Insert content into the document
    _, err = srv.Documents.BatchUpdate(createdDoc.DocumentId, &docs.BatchUpdateDocumentRequest{
        Requests: []*docs.Request{
            {InsertText: &docs.InsertTextRequest{
                Text: content,
                Location: &docs.Location{
                    Index: 1,
                },
            }},
        },
    }).Do()
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("https://docs.google.com/document/d/%s", createdDoc.DocumentId), nil
}
