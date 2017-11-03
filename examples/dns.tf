provider "ovh" {
    endpoint = "ovh-eu"
    application_key = "aaaaaaaaaaaaaa"
    application_secret = "aaaaaaaaa"
    consumer_key = "aaaaaaaaaaa"
}

resource "ovh_domain_zone_record" "test158" {
    zone = "testdemo.ovh"
    subdomain = "test158"
    fieldtype = "A"
    ttl = "3600"
    target = "0.0.0.0"
}
