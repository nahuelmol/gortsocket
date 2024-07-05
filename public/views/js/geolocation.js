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

var stopSharong = document.getElementById('stopsharingBtn');
var share = document.getElementById('shareBtn');
share.addEventListener('click', startSharing);
