package repositories

import (
	"context"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/memory"
	"github.com/objforce/objforce/cmd/index-api/app/domain/models"
	"github.com/objforce/objforce/cmd/index-api/app/providers"
	"testing"
	"time"
)

func TestCreateIndex(t *testing.T) {
	config, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
		return
	}

	data := []byte(`{
                "elastic": {
                        "urls": [
							"http://localhost:9200"
						]
                }
        }`)
	source := memory.NewSource(memory.WithJSON(data))

	err = config.Load(source)
	if err != nil {
		t.Fatal(err)
		return
	}

	client, err := providers.NewElasticClientProvider(config)
	if err != nil {
		t.Fatal(err)
	}

	indexRepo := NewDocumentRepository(client)


	doc := &models.Document{
		Index: "tweets",
		Type: "doc",
		Id: "1",
		Fields: map[string]interface{}{
			"name": "edison",
			"age": 18,
			"created_at": time.Now(),
		},
	}

	err = indexRepo.Upsert(context.Background(), doc)
	if err != nil {
		t.Fatal(err)
	}

	doc1, err := indexRepo.FindOne(context.Background(), doc.Index, doc.Type, doc.Id)
	if err != nil {
		t.Fatal(err)
	}
	if doc1 != nil {
		t.Log(doc1)
	}
}