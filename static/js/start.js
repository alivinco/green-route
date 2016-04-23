/**
 * Created by alivinco on 23/04/16.
 */



function loadData(){
    params = {ds:"datanorge",source:"lekeplasser",longitude:pos.lng,latitude:pos.lat,radius:1,date:"2015-08-04",school:"Auglend%20skole"}
    $.get( "/greenr/api/proxy",params, function( data ) {
      console.dir(data)
         $("#question_response").addClass("alert alert-success")
      $("#question_response").html("There are "+data.lekeplasser.length+" around you")
      console.dir(map)
      $.each(data.lekeplasser, function( index, value ) {
          //alert( index + ": " + value );
          setMarkerOnMap(parseFloat(value.latitude),parseFloat(value.longitude) ,"Play here")
      });
      //alert( "Load was performed." );
    });
}

function setMarkerOnMap(lat,lng,title)
{
    console.log("setting marker lat:"+lat+" lng:"+lng)
    //var map = new google.maps.Map(document.getElementById("map-canvas"), mapOptions);
    var myLatlng = new google.maps.LatLng(lat,lng);
    var marker = new google.maps.Marker({
        position: myLatlng,
        title:title,
        map:map
    });
    //marker.setMap(map);

}

