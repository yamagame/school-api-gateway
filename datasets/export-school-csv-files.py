import bpy
import csv
import os
from mathutils import Vector


def newfilepath(filepath, suffix):
    directory = os.path.dirname(filepath)
    basename = os.path.splitext(os.path.basename(filepath))[0]
    return os.path.join(directory, basename+suffix)


route_object = bpy.data.objects["Route"]

with open(newfilepath(bpy.data.filepath, "-edges.csv"), "w", newline='') as output:
    wcsv = csv.writer(output, delimiter=',')
    wcsv.writerow(["P0", "P1"])
    for v in route_object.data.edges:
        wcsv.writerow([v.vertices[0], v.vertices[1]])

with open(newfilepath(bpy.data.filepath, "-vertices.csv"), "w", newline='') as output:
    wcsv = csv.writer(output, delimiter=',')
    wcsv.writerow(["No", "X", "Y", "Z"])
    n = 0
    for v in route_object.data.vertices:
        x, y, z = v.co
        wcsv.writerow([n, x, y, z])
        n = n + 1

with open(newfilepath(bpy.data.filepath, "-objects.csv"), "w", newline='') as output:
    wcsv = csv.writer(output, delimiter=',')
    wcsv.writerow(["Name", "X", "Y", "Left", "Bottom", "Right", "Top"])
    for obj in bpy.data.collections["Buildings"].all_objects:
        l = obj.location
        b = [obj.matrix_world @ Vector(corner) for corner in obj.bound_box]
        wcsv.writerow([obj.name,
                       l[0], l[1],
                       b[0][0], b[0][1],
                       b[7][0], b[7][1]])
