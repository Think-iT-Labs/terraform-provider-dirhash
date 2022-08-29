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
