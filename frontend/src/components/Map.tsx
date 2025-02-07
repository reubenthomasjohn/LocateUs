import { useEffect, useRef } from "react";
import { Loader } from "@googlemaps/js-api-loader";
import { User } from "@/store/atoms/users";

interface MapProps {
  users: User[];
}

export function Map({ users }: MapProps) {
  const mapRef = useRef<HTMLDivElement>(null);
  const infoWindowRef = useRef<google.maps.InfoWindow | null>(null);

  useEffect(() => {
    const loader = new Loader({
      apiKey: "AIzaSyD-GiDN1eisMD78Qs7xEkAtQVgH1jCQInU",
      version: "weekly",
      libraries: ["visualization"],
    });

    loader.importLibrary("maps").then(() => {
      if (!mapRef.current) return;

      const map = new google.maps.Map(mapRef.current, {
        center: { lat: 18.597542302941, lng: 73.760751597583 },
        zoom: 14,
        styles: [
          {
            featureType: "all",
            elementType: "labels.text.fill",
            stylers: [{ color: "#6b7280" }],
          },
          {
            featureType: "water",
            elementType: "geometry.fill",
            stylers: [{ color: "#dbeafe" }],
          },
        ],
        mapId: "DEMO_MAP_ID",
      });

      // Create info window
      infoWindowRef.current = new google.maps.InfoWindow();

      console.log("Users: ", users);

      // Create markers for each location
      users.forEach((user) => {
        if (user.latitude && user.longitude) {
          const marker = new google.maps.Marker({
            position: { lat: user.latitude, lng: user.longitude },
            map: map,
            icon: {
              path: google.maps.SymbolPath.CIRCLE,
              scale: 7,
              fillColor: "#4F46E5",
              fillOpacity: 0.2,
              strokeWeight: 0.7,
              strokeColor: "#4338CA",
            },
          });

          // Add hover listener
          marker.addListener("mouseover", () => {
            if (infoWindowRef.current && user.latitude && user.longitude) {
              const content = `
                <div class="p-1">
                  <p class="font-semibold">${user.full_name}</p>
                  <p class="text-sm text-gray-600">${user.phone_number}</p>
                </div>
              `;
              infoWindowRef.current.setContent(content);
              infoWindowRef.current.open(map, marker);
            }
          });

          marker.addListener("mouseout", () => {
            if (infoWindowRef.current) {
              infoWindowRef.current.close();
            }
          });
        }
      });

      // Create heatmap layer
      const heatmapData = users.map(
        (user) => new google.maps.LatLng(user.latitude, user.longitude)
      );

      new google.maps.visualization.HeatmapLayer({
        data: heatmapData,
        map: map,
        radius: 25,
        opacity: 0.7,
      });
    });
  }, [users]);

  return <div ref={mapRef} className="w-full h-[400px] rounded-lg shadow-lg" />;
}
