
function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(sendPosition);
    } else {
        document.getElementById("location").innerText = "Geolocation is not supported by this browser.";
    }
}

function sendPosition(position) {
    var lat = position.coords.latitude;
    var lon = position.coords.longitude;

    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/bevisible?lat=" + lat + "&lon=" + lon, true);
    xhr.send();
    console.log("lat and lon")
}

function startSharing(){
    window.onload = getLocation;
    setInterval(getLocation, 1000);
}

function stopSharing(){
    console.log('this must stop the sharing process');
}

navigator.permissions.query({name: 'geolocation'}).then(function(permissionStatus) {
  // If permission was granted previously, fetch the location without asking again
  if (permissionStatus.state === 'granted') {
    getGeolocation();
  } else if (permissionStatus.state === 'prompt') {
    // Permission hasn't been granted or denied, request it now
    requestGeolocation();
  } else {
    console.log('Geolocation permission denied.');
  }

  permissionStatus.onchange = function() {
    console.log('Geolocation permission state has changed to ', this.state);
  };
});

var stopshare = document.getElementById('stopsharingBtn');
var share = document.getElementById('shareBtn');
share.addEventListener('click', startSharing);
stopshare.addEventListener('click', stopSharing);
