if (document.getElementById("link-app")){ 
  var linkApp = new Vue({
      delimiters: ['${', '}'],
      el: '#link-app',
      data: {
        url: "",
        text: "",
        code: "",
      },
      methods: {
        build: function (event) {
          if (this.url.length > 0 && this.text.length > 0) {
            this.code = "<a href=\"" + this.url + "\">" + this.text +"</a>"
          }
        }
      }
    })
}

if (document.getElementById("img-app")) {
  var imgApp = new Vue({
    delimiters: ['${', '}'],
    el: '#img-app',
    data: {
      url: "",
      code: "",
    },
    methods: {
      build: function (event) {
        if (this.url.length > 0) {
          this.code = "<img src=\"" + this.url + "\" />"
        }
      }
    }
  })
}