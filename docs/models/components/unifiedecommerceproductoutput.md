# UnifiedEcommerceProductOutput


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ProductURL`                                                                  | **string*                                                                     | :heavy_minus_sign:                                                            | The URL of the product                                                        |
| `ProductType`                                                                 | **string*                                                                     | :heavy_minus_sign:                                                            | The type of the product                                                       |
| `ProductStatus`                                                               | **string*                                                                     | :heavy_minus_sign:                                                            | The status of the product. Either ACTIVE, DRAFT OR ARCHIVED.                  |
| `ImagesUrls`                                                                  | []*string*                                                                    | :heavy_minus_sign:                                                            | The URLs of the product images                                                |
| `Description`                                                                 | **string*                                                                     | :heavy_minus_sign:                                                            | The description of the product                                                |
| `Vendor`                                                                      | **string*                                                                     | :heavy_minus_sign:                                                            | The vendor of the product                                                     |
| `Variants`                                                                    | [][components.Variants](../../models/components/variants.md)                  | :heavy_minus_sign:                                                            | The variants of the product                                                   |
| `Tags`                                                                        | []*string*                                                                    | :heavy_minus_sign:                                                            | The tags associated with the product                                          |
| `FieldMappings`                                                               | [*components.FieldMappings](../../models/components/fieldmappings.md)         | :heavy_minus_sign:                                                            | The custom field mappings of the object between the remote 3rd party & Panora |
| `ID`                                                                          | **string*                                                                     | :heavy_minus_sign:                                                            | The UUID of the product                                                       |
| `RemoteID`                                                                    | **string*                                                                     | :heavy_minus_sign:                                                            | The remote ID of the product in the context of the 3rd Party                  |
| `RemoteData`                                                                  | [*components.RemoteData](../../models/components/remotedata.md)               | :heavy_minus_sign:                                                            | The remote data of the customer in the context of the 3rd Party               |
| `CreatedAt`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | The created date of the object                                                |
| `ModifiedAt`                                                                  | **string*                                                                     | :heavy_minus_sign:                                                            | The modified date of the object                                               |