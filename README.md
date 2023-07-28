# Groupie Tracker Geolocalization

The Groupie Tracker Geolocalization project is a web application that allows users to map the different concert locations of a certain artist or band provided by the Client. The backend of the application is written in Go, and it utilizes a process of converting addresses into geographic coordinates to place markers for the concerts on a map. 

## Objectives

The main objectives of the Groupie Tracker Geolocalization project are as follows:

1. Convert addresses (e.g., "Germany Mainz") into geographic coordinates (e.g., 49.59380, 8.15052).
2. Display markers on the map representing the concert locations of a certain artist or band.
3. Create a user-friendly web interface to interact with the application.
4. Handle website errors gracefully to provide a smooth user experience.


## Technologies Used

The project utilizes the following technologies:

- Backend: Go
- Frontend: HTML, CSS, JavaScript
- Maps API: Google Maps API
- Geocoding Service: Google Maps Geocoding API
- Data Storage: JSON files


## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@git.01.alem.school:mnazarbe/groupie-tracker-geolocalization.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd groupie-tracker-geolocalization
```
Run a program:
```CMD/Terminal 
go run main.go
```

Follow the link on the terminal:
```CMD/Terminal 
Server is listening...http://127.0.0.1:7777
```

## Conclusion

The Groupie Tracker Geolocalization project provides an interactive and user-friendly way to visualize the concert locations of a certain artist or band. It demonstrates data manipulation, geolocation, geocoding, and interaction with Maps APIs. The application is written in Go, follows good coding practices, and includes unit tests for reliability. Enjoy tracking your favorite artists' concerts on the map!

- azakonov
- mnazarbe