data "external_schema" "gorm" {
    program = [
        "go",
        "run",
        "-mod=mod",
        "./cmd/migrations/main.go"
    ]
}

env "gorm" {
    src = data.external_schema.gorm.url
    dev = "docker://postgres/16/dev?search_path=public"
    migration {
        dir = "file://migrations"
    }
    format {
        migrate {
            diff = "{{ sql . \"  \" }}"
        }
    }
}
