import math

def calculate_area(min_lat, max_lat, min_long, max_long):
    # Convert latitude and longitude to radians
    min_lat_rad = math.radians(min_lat)
    max_lat_rad = math.radians(max_lat)
    min_long_rad = math.radians(min_long)
    max_long_rad = math.radians(max_long)

    # Calculate average latitude
    average_lat = (min_lat_rad + max_lat_rad) / 2.0

    # Calculate the conversion factor for longitude based on average latitude
    lon_conversion = math.cos(average_lat)

    # Calculate the area in square meters
    area = (max_lat_rad - min_lat_rad) * 111319.9 * (max_long_rad - min_long_rad) * lon_conversion

    return area

# Given coordinates
min_lat = 59.17622388212263
max_lat = 59.209588091683415
min_long = 17.58278807473854
max_long = 17.663812244660416

# Calculate and print the area
area_square_meters = calculate_area(min_lat, max_lat, min_long, max_long)
print(f"The approximate area is: {area_square_meters:.2f} square meters")
