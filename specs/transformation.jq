  # Fix the file downloads: invalid swagger for file downloads
  # Rename msaspec.Error back to msa.APIError. These are two names for the same type.
  walk(
    if type == "object" and has("$ref") and ."$ref" == "#/definitions/msaspec.Error" then ."$ref" = "#/definitions/msa.APIError" else . end
    | if type == "object" and has("$ref") and ."$ref" == "#/definitions/msaspec.MetaInfo" then ."$ref" = "#/definitions/msa.MetaInfo" else . end
    )
  | del(.definitions."msaspec.Error")

  # Misc fixes
  | .paths."/oauth2/token".post.responses."201".headers["X-CS-Region"] = {type: "string"}