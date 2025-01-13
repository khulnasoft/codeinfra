resource "prov" "codeinfra:providers:simple" {

}

resource "res" "simple:index:Resource" {
    options {
        provider = prov
    }
    
    value = true
}