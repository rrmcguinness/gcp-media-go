load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
        version = "v0.0.0-20190523083050-ea95bdfd59fc",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:WBZRG4aNOuI15bLRrCgN8fCq8E5Xuty6jGbmSNEvSsU=",
        version = "v0.0.0-20191209042840-269d4d468f6f",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:ga8SEFjZ60pxLcmhnThWgvH2wg8376yUJmPhEH4H3kw=",
        version = "v0.0.0-20240423153145-555b57ec207b",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:IgJPqnrlY2Mr4pYB6oaMKvFvwJ9H+X6CCY5x1vCTcpc=",
        version = "v0.12.1-0.20240621013728-1eb8caab5155",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:gVPz/FMfvh57HdSJQyvBtF00j8JU4zdyUgIUNhlgg0A=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:6pFjapn8bFcIbiKo3XT4j/BhANplGihG6tvd+8rYgrY=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_yaml_yaml",
        importpath = "github.com/go-yaml/yaml",
        sum = "h1:RYi2hDdss1u4YE7GwixGzWwVo47T8UQwnTLB6vQiq+o=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:OptwRhECazUx5ix5TTWC3EZhsZEHWcYWY4FQHTIubm4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_google_generative_ai_go",
        importpath = "github.com/google/generative-ai-go",
        sum = "h1:6ybg9vOCLcI/UpBBYXOTVgvKmcUKFRNj+2Cj3GnebSo=",
        version = "v0.18.0",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:PVRnTgtArZ3QQqTGtbtjtnIkzl2iY2kt24yqbrf7td8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:DIhPTQrbPkgs2yJYdXU/eNACCG5DVQjySNRNlflZ9Fc=",
        version = "v3.3.3",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:XYIDZApgAnrN1c855gTgghdIA6Stxb52D5RnLI1SLyw=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:8gw9KZK8TiVKB6q3zHY3SBzLnrGp6HQjyfYBYGmXdxA=",
        version = "v2.12.5",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_pkg_diff",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:GFCKgmp0tecUJ0sJuv4pzYCqS9+RGSn52M3FUwPs+uo=",
        version = "v0.6.1-0.20240319094008-0393e58bdf10",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:CnFSK6Xo3lDYRoBKEcAtia6VSC837/ZkJuRduSFnr14=",
        version = "v0.115.0",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:vO95gvBi7qUgfA9SflexQs9hB4U4tnri/GwADIrLQy8=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:GgdNoDwZR5RIO3j8XwXqa6Gc6q5mP3KYMdFC7FEVyG4=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_ai",
        importpath = "cloud.google.com/go/ai",
        sum = "h1:rXUEz8Wp2OlrM8r1bfmpF2+VKqc1VJpafE3HgzRnD/w=",
        version = "v0.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:EPPqgHDJpBZKRvv+OsB3cr0jYz3EL2pZ+802rBPcG8U=",
        version = "v1.68.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:O0fj88npvQFxg8LfXo7fArcSrC/wtAstGuWQ7dCHWjg=",
        version = "v0.23.2",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:DO5Vn3zmY1aDyfoqni8e8+x+lwrfLCoAAbEui9NB0y8=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:z08Xuv7ZtaB2d4jsJi9/WhbnnI5s19wlLDZpssn3Fus=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:o1C/+IvzwYeV1doum61XmJQ/Bwpk/4+2DT1JyVu2x64=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:qYrjEHEFY7+CL4QlHIHuwTgrTnZbSKzdPFqgjZDsQNo=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:sUrR96yokdL6tTTXK0X13V1TLMta8/1u328bRG5lWZc=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:SSvoD0ofOydm5gA1++15pW9VPgQbk0OmNlcb7JczoO4=",
        version = "v1.14.9",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:mCqyoaDjDzaW1RqmmQtCJuawb9nca5bEu7HvVcpZDwg=",
        version = "v1.19.1",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:xieyFA+JKyTDkO/Z9UyVEpkHW8pDYykU51O4G0pvXEg=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_auth",
        importpath = "cloud.google.com/go/auth",
        sum = "h1:uiha352VrCDMXg+yoBtaD0tUF4Kv9vrtrWPYXwutnDE=",
        version = "v0.7.2",
    )
    go_repository(
        name = "com_google_cloud_go_auth_oauth2adapt",
        importpath = "cloud.google.com/go/auth/oauth2adapt",
        sum = "h1:0GWE/FUsXhf6C+jAkWgYm7X9tK8cuEIfy19DBn6B6bY=",
        version = "v0.2.4",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:w9AyogtMLXbcy5kzXPvk/Q3MGQkgJH7ZDB8fAUUxTt8=",
        version = "v1.13.7",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:W4oSMS6vRCo9DLr1RPyDP8oeLverbvhJRzaZSsipft8=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:zaQwOAd7TlE84pwPHavNMsnv5zRyRV8ym2DJ4iQ2cV0=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:KBcujO3QRvBIwzZLtvQEPB9SXdovHnMBx0V/uhucH9o=",
        version = "v1.0.6",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:w2Goy9n6gh91LVi6B2Sc+HpBl8WbWhIyzdvVvrAuEIw=",
        version = "v1.61.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:GbOg1uGvoV8FXxMStFoNcq5z9AEUwCpKt/6GNcuDSZM=",
        version = "v1.18.5",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:RHnEM4HXbWShlGhPA0Jzj2YYETCHxmisNMU0OE2fXQM=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:XURrQhj5COWAEvICivbGID/Hu67AvMYHAhMRIyc3Ux8=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:PrplNaAS6Dn187e+OcGzyEKETX8iL3tCaDqcPPW7Zoo=",
        version = "v1.17.7",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:zkCG1dBezxRM3dtgQ9h1Y+IJ7V+lARWgp0l9k/SZsfU=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:Q47KKoA0zsNcC9U5aCmop5TPPItVq4cx7Wwqgra+5PU=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:Y0HUuiCAVk9BojLItOycBl91tY25NXH8oFsyi1IC/U4=",
        version = "v1.12.8",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:EGawh2RUnfHT5g8f/FX3Ds6KZuIBC77hZoDrBvEZw94=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:NM6oZeZNlYjiwYje+sYFjEpP0Q0zCan1bmQW/KmIrGs=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:46ertIh+cGkTg/lN7fN+TOx09SoM65dpdUp96vXBcMY=",
        version = "v1.13.2",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:KZ5IbxnUNcTdBkMfZNum6z4tu61sFHf9Zk7Xl8RzWVI=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:mSrneOVadcpnDZHJebg+ts/10azGTUKOCSQET7KdT7g=",
        version = "v0.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:czcba5mxwRM5V//jSadyig0y+8aOHmN7gUl9GbHu59E=",
        version = "v1.20.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:wKEakCbRevlwsWqTn34pWJUFmdbx0HKwpRH6HhU7NIs=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:MiK1Us7YP9+sdNViUE4X2B2vLScrKcjOPw5b6uamZvE=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:ViFnMnUK7LNcWvisZgihxXit76JxSHFeijYI5U/gjOE=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:M6irSHns6VxMro+IbvDxDJLD6tkfjlW+mo2MPaM23KA=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:e8SV0yKuSjgHEZaQcZwjKXe0ta1jZrvLxX/2i/IAG+8=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:RNMG5ffWKdbWOkwvjC4GqxLaxEaWFpm2hQCF2WFW/vo=",
        version = "v2.4.2",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:qM60MGNTGsSJuzAziVJjtRA7pGby2dA8OuqdVRe/lYo=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:6Me8ugrAOAxssGhSo8im0YSuy4YvYk4mbGvCadAH5aE=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:FfNUy9j3aRQ99L4a5Rdm82RMuiw0BIe3lpPn2ykom8k=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:fzbObuGgoViO0ArFuOQIJ2yr5bH5YzbORVvMDBrDC5I=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:DgCHYpasCaI3b03X6pxqGmEzX72DAHZhV0f/orrpxYM=",
        version = "v1.54.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:/GQVl5gOPR2dUemrR2YJxZG5D9MCE3AYgmDxjzP54jI=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:6KI6P04WExzrfbciW5RTEQScBEY98Fc4VtS04ufT3Js=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:IixFIMRzUJWZUAOe8s/K2X4Bvtp0A3xjHLljfNC4aSo=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:xa6MIQhGylE24QdWaxhfIfAJE3Pupcr+i77WEx3NJrg=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:p5Y7ZNVPiV9pEAHzvWiPcSiQRMQqcuHxOP0ZOP0vVww=",
        version = "v1.6.8",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:we+qx5uCZ88aQzQS3MJXRvAh/ik+EmqVyjcW1oYFW44=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:CpRnsUpMU5gxUKyfh7TD0SM+E+7E4ORaDea2JctKfpY=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:/k8ppuWOtNuDHt2tsRV42yI21uaGnKDEQnRFeBpbFF8=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:83bd2lCgtu2nLbX2jrqsrQhIs7VuVA1N6Op5syeRVIg=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:wysUXEkggPwENZY3BXroOyWoyVfPypzaqNHgOZD9Kck=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:BfXsTXYs5xlicAlgbtlo8Cw+YdzU3PrlBg7dATJUwrk=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:bHwcvgh8AmcYm6p6/ZrWW3a7J7sKBDtqtsyVXKssnPs=",
        version = "v0.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:zaWBakKPT6mPHVn5iefuRqttjpbNsb8LlMw9KgfyfyU=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:06Jg3JeLslEfBYX1sDqOPLnF7a3wmhNcDUXF/fVOb50=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:r7umDwhj+BQyz0ScZMp4QrGXjSTI3ZINnpgU2nlB/K0=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:rcuRS9XfOgr1v6TAoihVeSXntOnpVhFlVHtPfgOkLAo=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:wtd+r415yrfZ8LsB6yH6WrOZ26tYt7w6wy3i5a4HQZ8=",
        version = "v1.4.7",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:M9SKIj9eoxoXCzytkLZVAuf5wmoui1OeDqEjC97wRbY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:5k0wXqkxL+YcXd4viQzTqCgzzVKKxzgrK+rCZJytEQs=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:kOYJEcuZgyUX/i/4DFrfXPcrddm1XCQD2lDI5hIFmZQ=",
        version = "v1.12.5",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:qqEmApr5YFOQjkrU8Jy6o6QpkESqfGbfrE6bnUZZbV8=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:f+ZXMqyrSJ5vZ5pE/zr0xC8y/M9BLNzQeLBwfeZ+wY4=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:WLbHekDbjK1fVFD3ibpFFVoyizlLRl73I7YKuAKilhU=",
        version = "v0.5.7",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:uWA9WQyfA0JdkeAFymWUsa3qE9tC33LUElla790Ou1A=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:2U1NB/GIoXhNNmYMlGiLM7juL7nxh51lSkNELbYddB8=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:izgww3TlyvWyDWdFKnrASpbh12IkAuw8o2ION8sAjX0=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:hE7f3ze3+eWh/EbYXEz7oXkm0LXcr7UCoLklwi7gsLU=",
        version = "v1.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:otHcJkci5f/sNRedrSM7eM81QRnu0yZ3HvkvWGphABA=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:NCXf8hfQi+Kmr56QJezXRZ6GPb80ZI7El1XztyUuLQI=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:jYpQ86mZ7OYZc7WadvCIlIaPXmXhr5nD7wgE/ekMVpM=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:Ex1/aYkA0areleSmOGXHvEFBGohteIYJr2SGPrjOUe0=",
        version = "v1.13.2",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:aepEkfiwOvUL9eu3ginVZhTaXDRHncQKi9lTT1BycH0=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:sFU1ETg1HfIN/Tev8gD0dleAITLv7cHp0JClwFmJ6bo=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:FPfowA/LEckKTQT0A4NJMI2bSou999c2ZyFX1zGiYxY=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:C2WL4ZnclXsh4XickGhKYKlPjqVZj35y1sbRjdsZ3g4=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:fGftW2bPi8vTjQm57xlwtLBZQcrgC+c3HMFBzJ+KWPc=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:HXsXGFaFaLTklwKgSob/GSE+c3verYDQDgreFaosxyc=",
        version = "v1.12.7",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:7AgOWH1oMPrB1AVU0/f47ADdOt+XfdBY7QRb8tcMUp8=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:CbCjfR/pgDHyRMu94o9nuGwaONEcarWnUfSGGw+I2ZI=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:LGt85MZUKlq9oqsbBL9+M6jAyeuR1TtCx6k5HfAQxTY=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:wGZKKJhYyuf4gcAEywQqQ6F19yxhBJGnzgyxOTbJjBw=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:J1OT7h51ifATIedjqk/uBNPh+1hkvUaH4VKbz4UuAsc=",
        version = "v1.38.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:jLQozsEVr+c6tOU13vDugtnaBSUy/PD5zK6mhm+uF1Y=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:+QG02kE63W13vXI+rwAxFF3EhGX6K7gXwFz9OKwKcHw=",
        version = "v2.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:N6n/TEr0FQzeP4ZtvF5daMszOhdZI94uMiPiAi9kFMo=",
        version = "v0.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:v9x75vXP5wMXw3QG3xmgjVHLlqYufuLn/ht3oNWCA3w=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:1veL/h/x5bgzG2CLK2cdG3plWdgO0p1qoHgwFBqG7+c=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:SdvD0PaPX60+yeKoSe16mawFpM0EPuiPPihTIVlhRsY=",
        version = "v1.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:yEuByg5XBHhTG9wPEU7GiEtC9Orp1wSEyiiX4IPqoSY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:YTKwc6K02xpa/SYkPpxY7QEmsd3deP/+ceMTuAQ1RVg=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:E4Z5e681Qh7UJrJRMCgYhp+3tkcoXiaKGh3UZmUPaAQ=",
        version = "v1.3.7",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:Jn/unfNUgRiNJRc1nrApzimKiVj91UYlLT8mMfpUu48=",
        version = "v1.10.8",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:TTGo2Vz7ZxYn2QbmuFP7Zo4lDm5VsbzBjDReo3SA5h4=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:u4RCnEQPvlrrnFRFinU0T3WsjtrsQErkWBfqTM5oUQI=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:Y8C0I/mzLbaxAl5cw3EaLox0Rvpy+VUwEuCGWIQDMU8=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:c3OAhTcZ8LbIiKps5T3p6i0QcPI8/aWYwOfoZobICKo=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:HxCzcUxSsCh6FJWkmbOUrGI1sKe4E1Yy4vaykn4RhJ4=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:P6+BY70Wtol4MtryBgnXZVTZfsdySEvWfz0EpyLwHi4=",
        version = "v1.63.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:TcWEAOLQH1Lb2fhHS6/GjvAh+ue0dt4xUDHXHG6vF04=",
        version = "v1.23.1",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:RusiwatSu6lHeEXe3kglxakAmAbfV+rhtPqA6i8RBx0=",
        version = "v1.41.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:CXmoNEvz7y2NtHFZuH3Z8ASN43rxRINWa2Q/IlBzM2k=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:RoyEtftfJrbwJcu63zuWE4IjC76xMyVsJBhmleIp3bE=",
        version = "v1.6.8",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:qR6Mu+EM2OfaZR1/Rl8BDBTVfi2X5OtwKKvJRSQyG+o=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:ngQokxUB1z2gvHn3vAf04m7SFnNYMiQIIpny81fCGAs=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:gK8z2BIJQ3KIYGddw9RJLne5Fx0FEXkrEQzPaeEYVvk=",
        version = "v1.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:g+B29z4gtRGsiKDoTF+bNeH25bLRokAaElygX2FcZkE=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:ue/1C8TF8H2TMzKMBdNnFxT7QaeWMtqfDr9TSQGgUhA=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:SKBkFTuOclESLjQL1LwraqVFm2fL5oL9tbzKITU+FOY=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:j9RxG8DcyJO/D7/ps2pOey8VZys+TMqF79bWAhuM7QU=",
        version = "v2.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:bf2qKqEN7iqT62IptQ/FDadoDLJI9sthyrW3PVaH8bY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:x4KwHB4JlBEzMaITVhrbbpHrU+2I5LrlvHGEEluT0vc=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:F5woMLufKnshmDvPVxCzoC+Di12RYXQ1W8kNmpBT8z0=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:EWTSVagWWeQjVAsebiF/wJMwC5bq6Zz3LqOmD9Uid4s=",
        version = "v1.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:R5OW5SNRqD0DSEmyWLUMNYAXWYnz/NLSXBawVFrc9a0=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:2bE69mh68law1UZWPjgmvOQsjsGSppRudABAXwNAy58=",
        version = "v1.12.6",
    )
    go_repository(
        name = "dev_cel_expr",
        importpath = "cel.dev/expr",
        sum = "h1:O1jzfJCQBfL5BFoYktaxwIhuttaQPsVWerH9/EEKx0w=",
        version = "v0.15.0",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:r6I7RJCN86bpD/FQwedZ0vSixDpwuWREjW9oRMsmqDc=",
        version = "v0.54.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:TT4fX+nBOA/+LUkobKGW1ydGcn+G3vRw9+g5HwCphpk=",
        version = "v0.54.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:PdomN/Al4q/lN6iBJEN3AwPvUiHPMlt93c8bqTG5Llw=",
        version = "v1.29.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:vPf/HFWTNkPu1aYeIsc98l4ktOQaL6LeSoeV2g+8YLc=",
        version = "v1.29.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:vkqKjk7gwhS8VaWb0POZKmIEDimRCMsopNYnriHyryo=",
        version = "v1.29.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:J/8ZNK4XgR7a21DZUAsbF8pZ5Jcw1VhACmnYt39JTi4=",
        version = "v1.29.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:n2OPp+PPXX0Axh4GuSsL5QL8xQCTb2oDwyzPnQvqUug=",
        version = "v0.186.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
        version = "v1.4.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:CUiCqkPw1nNrNQzCCG4WA65m0nAmQiwXHpub3dNyruU=",
        version = "v0.0.0-20240617180043-68d350f18fd4",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:hjSy6tcFQZ171igDaN5QHOw2n6vx40juYbC/x67CEhc=",
        version = "v0.0.0-20240903143218-8af14fe29dc1",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:Rie8vnNXn/RjOgFacUrolQKaHsN10UPAXBb3IkfDdE4=",
        version = "v0.0.0-20240617180043-68d350f18fd4",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:pPJltXNxVzT4pK9yD8vR9X75DaWYYmLGMsEvBfFQZzQ=",
        version = "v0.0.0-20240903143218-8af14fe29dc1",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:3QdXkuq3Bkh7w+ywLdLvM56cmGvQHUMZpiCzt6Rqaoo=",
        version = "v1.66.2",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:6xV6lTsCfpGD21XK49h7MhtcApnLqkfYgPcdHftf6hg=",
        version = "v1.34.2",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:GXm2NjJrPaiv/h1tb2UH8QfgC/hOf/+z0p6PT8o1w7A=",
        version = "v0.27.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
        version = "v0.0.0-20190121172915-509febef88a4",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
        version = "v0.0.0-20190313153728-d0100b6bd8b3",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:5ORfpBpCs4HzDYoodCDBbwHzdR5UrLBZ3sOnUJmFoHo=",
        version = "v0.29.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:PbgcYx2W7i4LvjJWEbf0ngHV6qJYr86PkAV3bXdLEbs=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:3NFvSEYkUoMifnESzZl15y791HH1qU2xm6eCJU5ZPXQ=",
        version = "v0.8.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:r+8e+loiHxRqhXVl6ML1nO3l1+oFoWbnlu2Ehimmi34=",
        version = "v0.25.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:Mh5cbb+Zk2hqqXNO7S1iTjEphVL+jb8ZWaqh/g+JWkM=",
        version = "v0.24.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:XvMDiNzPAl0jr17s6W9lcaIhGUfUORdGCNsuLmPG224=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:eTDhh4ZXt5Qf0augr54TN6suAUudPcawVZeIAPU7D4U=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:vU5i/LfpvrRCpgM/VPfJLg5KjxD3E+hfT1SH+d9zLwg=",
        version = "v0.21.1-0.20240508182429-e35e4ccd0d2d",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:+cNy6SZtPcJQH3LJVLOSmiC7MMxXNOb3PU/VUEz+EhU=",
        version = "v0.0.0-20231012003039-104605ab7028",
    )
    go_repository(
        name = "tech_einride_go_aip",
        importpath = "go.einride.tech/aip",
        sum = "h1:d/4TW92OxXBngkSOwWS2CH5rez869KpKMaN44mdxkFI=",
        version = "v0.67.1",
    )

GO_TEST_TYPES = [
    "@com_github_stretchr_testify//assert",
]
