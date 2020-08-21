<template>
    <div class="otaks-map">
        <v-img height="100%" width="100%">
        <l-map 
            style="height: 100%; width: 100%"
            :zoom="zoom"
            :center="center"
        >
            <l-control-scale position="topright" :imperial="true" :metric="false"></l-control-scale>

            <l-tile-layer :url="url"></l-tile-layer>
            <!-- <l-marker 
                :lat-lng="markerLatLng"
                :icon="atakUserIcon">
            </l-marker> -->
            <l-marker :lat-lng="markerLatLng">
                <l-icon
                    :icon-anchor="dynamicAnchor"
                    :icon-size="dynamicSize"
                    :icon-url="atakUserSymbolUrl"
                >
                </l-icon>
            </l-marker>
        </l-map>
        </v-img>
    </div>
</template>

<script>
import { 
    LMap, 
    LTileLayer, 
    LMarker, 
    LIcon, 
    LControlScale 
} from 'vue2-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { ms, std2525c }from 'milsymbol/index.esm.js'

ms.addIcons(std2525c);

var atakUserSymbol = new ms.Symbol("SFGPU-------", {
    size: 5,
    quantity: 1
});
var atakUserSymbolUrl = atakUserSymbol.toDataURL()
console.log("atakUserSymbol::toDataURL()", atakUserSymbolUrl);

delete L.Icon.Default.prototype._getIconUrl;

L.Icon.Default.mergeOptions({
   iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
   iconUrl: require('leaflet/dist/images/marker-icon.png'),
   shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
});

// var atakUserIcon = L.Icon({
//     iconUrl: atakUserSymbolUrl,
//     iconSize: [5, 5],
//     iconAnchor: [atakUserSymbol.getAnchor().x, atakUserSymbol.getAnchor().y],
// });


export default {
    name: "otaks-map",
    components: {
        LMap,
        LTileLayer,
        LMarker,
        LIcon,
        LControlScale,
    },
    data: () => ({
        url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
        zoom: 7,
        center: [35, -115],
        bounds: null,
        markerLatLng: [37.245056, -115.816548],
        //markerLatLng: [47.313220, -1.319482],
        atakUserIcon: new L.Icon({
            iconUrl: atakUserSymbolUrl,
            iconSize: [25, 25],
            iconAnchor: [atakUserSymbol.getAnchor().x, atakUserSymbol.getAnchor().y],
        }),
        atakUserSymbolUrl: atakUserSymbolUrl,
    }),
    computed: {
        dynamicSize () {
            return [this.iconSize, this.iconSize * 1.15];
        },
        dynamicAnchor () {
          return [this.iconSize / 2, this.iconSize * 1.15];
        }
    }
}
</script>

<style>
.otaks-map {
    height: 100%;
    width: 100%;
    display: inline;
}
</style>