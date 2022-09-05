# Dirhash
Computing the checksum of a directory made easy.

## Dirhash Provider
Dirhash was created to compute the checksum of a directory, 
a useful provider if you try to make Terraform react and make changes 
based on whether changes have been made inside a directory or not.


### Usage
As noted above, the Dirhash provider can compute the checksum of a directory, 
the data_source `dirhash_sha256` can be used to retrieve the SHA256 of a directory.
For example:
```terraform
provider "dirhash" {}
```
```terraform
data "dirhash_sha256" "example" {
  directory = "/path/to/directory"
  ignore = [
    "glob_pattern_1/*",
    "glob_pattern_2/*"
  ]
}
output "directory_sha256_checksum" {
  description = "the Output SHA256 Checksum for the directory"
  value       = data.dirhash_sha256.example.checksum
}
```
