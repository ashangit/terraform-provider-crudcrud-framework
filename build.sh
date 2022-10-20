rm -rf /Users/nicolas.fraison/Documents/dev/github/terraform-provider-crudcrud-framework/examples/.terraform*

rm -rf /Users/nicolas.fraison/Documents/dev/github/terraform-provider-crudcrud-framework/examples/terraform.tfstate*

mkdir -p ~/.terraform.d/plugins/local/axaclimate/crudcrud/0.1.0/darwin_arm64

go build -o bin/terraform-provider-crudcrud_v0.1.0

mv bin/terraform-provider-crudcrud_v0.1.0 ~/.terraform.d/plugins/local/axaclimate/crudcrud/0.1.0/darwin_arm64
