# Things To Do Generator

A generator of things to do given a location that ranks things based on Yelp reviews.

"Itineraries" include:
- activity
- food
- +1 low maintenance (some activity that does not take a lot of time)

FR:
- As a user I should be able to give a location and a day and have a list of generated itineraries from yelp reviews

NFR:
- none really I am just coding this for myself who cares how scalable (pr lack thereof) it is

Yelp Integration
- uses the /businesses/search endpoint to look for things, takes:
    - term: query
    - location: city or coord
    - price: price
    - sort_by: usually by rating or review_counts (probably a mixture of the 2 or we might have to do it on our end)

