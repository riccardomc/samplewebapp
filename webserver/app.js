var url = 'http://localhost:8080/data/';

vm = new Vue({
  el: '#app',
  data: {
    messages: [
      'a',
      'b',
      'c'
    ]
  },
  methods:{
    getData: function() {
      fetch(url).then(function(response){
        return response.json();
      }).then(function(data) {
        vm.messages = data;
      }).catch(function() {
        console.log("Fetch failed: " + url);
      });
    }
  },
  mounted: function() {
    this.getData();
  }
});