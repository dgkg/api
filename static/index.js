// Initialize and add the map
function initMap() {
  // The location of Uluru
  const toulouse = { lat: 43.6046, lng: 1.4443 };
  // The map, centered at toulouse
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 15,
    center: toulouse,
  });
  
  const womenImg = "static/women.png";
  new google.maps.Marker({
    position: toulouse,
    map: map,
    icon: womenImg,
  });

  const toulouseM = { lat: 43.6049, lng: 1.4446 };
  const manImg = "static/man.png";
  new google.maps.Marker({
    position: toulouseM,
    map: map,
    icon: manImg,
  });

  map.setOptions({ styles: styles["hide"] });

  // The marker, positioned at toulouse
  // const marker = new google.maps.Marker({
  //   position: toulouse,
  //   map: map,
  // });

  //const t1 = { lat: 43.604652, lng: 1.444208 };

  // new google.maps.Marker({
  //   position: t1,
  //   map: map,
  // });

  //const markerCluster = new markerClusterer.MarkerClusterer({ map, markers });

  //NewMarker(map, toulouse);

  const toulouse1 = { lat: 43.6067, lng: 1.444 };

  const taxiImg = "static/taxi.png";
  new google.maps.Marker({
    position: toulouse1,
    map,
    icon: taxiImg,
  });

  const toulouse2 = { lat: 43.6077, lng: 1.443 };

  new google.maps.Marker({
    position: toulouse2,
    map,
    icon: taxiImg,
  });
}

// create marker using svg icon
// function NewMarker(map,position) {
//   // change color if this cluster has more markers than the mean cluster
//   const color = 10 > Math.max(10, 100) ? "#ff0000" : "#0000ff";

//   // create svg url with fill color
//   const svg = window.btoa(`
// <svg fill="${color}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 240">
// <circle cx="120" cy="120" opacity=".6" r="70" />
// <circle cx="120" cy="120" opacity=".3" r="90" />
// <circle cx="120" cy="120" opacity=".2" r="110" />
// <circle cx="120" cy="120" opacity=".1" r="130" />
// </svg>`);

//   return new google.maps.Marker({
//     position,
//     icon: {
//       url: `data:image/svg+xml;base64,${svg}`,
//       scaledSize: new google.maps.Size(45, 45),
//     },
//     label: {
//       text: String(10),
//       color: "rgba(255,255,255,0.9)",
//       fontSize: "12px",
//     },
//     // adjust zIndex to be above other markers
//     zIndex: 1000 + 10,
//   });
// }

const styles = {
  default: [],
  hide: [
    {
      featureType: "poi",
      stylers: [{ visibility: "off" }],
    },
    {
      featureType: "transit",
      elementType: "labels.icon",
      stylers: [{ visibility: "off" }],
    },
  ],
};

// [
//   {
//     "featureType": "all",
//     "stylers": [
//       { "color": "#C0C0C0" }
//     ]
//   },{
//     "featureType": "road.arterial",
//     "elementType": "geometry",
//     "stylers": [
//       { "color": "#CCFFFF" }
//     ]
//   },{
//     "featureType": "landscape",
//     "elementType": "labels",
//     "stylers": [
//       { "visibility": "off" }
//     ]
//   }
// ]

window.initMap = initMap;
