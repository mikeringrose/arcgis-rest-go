# arcgis-rest-go

## Example Usage
```
import (
	"context"
	"fmt"

	"mikeringrose.com/esri/arcgis"

	"golang.org/x/oauth2"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := arcgis.NewClient(tc, "https://www.arcgis.com/sharing/rest")
	self, _ := client.Portal.Self()
	fmt.Printf("Self: %+v", self)
}
```
