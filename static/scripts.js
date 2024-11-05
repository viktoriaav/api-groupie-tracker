function togglePopup(id){
    var popup = document.getElementById(id);
    popup.classList.toggle("active");
};

let artists=document.getElementById('ArtistScrollBox')
    artists.addEventListener('wheel',(event)=>{ 
        event.preventDefault();
        artists.scrollLeft += event.deltaY;
});

$("button").click(function(){
    $.get("demo_test.asp", function(data, status){
      alert("Data: " + data + "\nStatus: " + status);
    });
  });