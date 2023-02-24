package prefect2

import (
	"context"
	"log"
	"strconv"
	"time"

	hc "terraform-provider-prefect2/prefect2_api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBlockSchemas() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBlockSchemasRead,
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"block_schema_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"checksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"block_schemas": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"block_type_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBlockSchemasRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspaceId := d.Get("workspace_id").(string)
	blockSchemaId := d.Get("block_schema_id").(string)
	checksum := d.Get("checksum").(string)

	log.Printf("MAAIKEL")
	log.Printf(checksum)

	var blockSchemasOutput []interface{}

	if blockSchemaId != "" {
		log.Printf("MAAIKEL 1")
		blockSchema, err := c.GetBlockSchemaById(blockSchemaId, workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}

		blockSchemas := make([]hc.BlockSchema, 1, 1)
		blockSchemas[0] = *blockSchema
		blockSchemasOutput = tfBlockSchemaSchemaOutput(blockSchemas)
	} else if checksum != "" {
		log.Printf("MAAIKEL 2")
		blockSchema, err := c.GetBlockSchemaByChecksum(checksum, workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}

		blockSchemas := make([]hc.BlockSchema, 1, 1)
		blockSchemas[0] = *blockSchema
		blockSchemasOutput = tfBlockSchemaSchemaOutput(blockSchemas)
	} else {
		log.Printf("MAAIKEL 3")
		blockSchemas, err := c.GetAllBlockSchemas(workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}
		blockSchemasOutput = tfBlockSchemaSchemaOutput(blockSchemas)
	}

	log.Printf("%+v\n", blockSchemasOutput)

	if err := d.Set("block_schemas", blockSchemasOutput); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	log.Printf("MAIKEL D")
	log.Printf("%+v\n", d)

	return diags
}

func tfBlockSchemaSchemaOutput(block_schemas []hc.BlockSchema) []interface{} {
	schemaOutput := make([]interface{}, len(block_schemas), len(block_schemas))

	for i, block_schema := range block_schemas {
		schema := make(map[string]interface{})

		schema["id"] = block_schema.Id
		schema["block_type_id"] = block_schema.BlockTypeId

		schemaOutput[i] = schema
	}
	return schemaOutput
}
