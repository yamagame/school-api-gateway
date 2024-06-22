import bpy
import csv
import os
import bmesh
import math


def newfilepath(filepath, suffix):
    directory = os.path.dirname(filepath)
    basename = os.path.splitext(os.path.basename(filepath))[0]
    return os.path.join(directory, basename+suffix)


object_name = "RouteFromCSV"

verts = []
edges = []
faces = []

# read file change file path
csvPoints = newfilepath(bpy.data.filepath, "-route.csv")
pointsReader = csv.reader(open(csvPoints, newline=''), delimiter=',')

# add xyz to verts
for row in pointsReader:
    vert = (float(row[0]), float(row[1]), 0)
    verts.append(vert)

# create mesh and object
mesh = bpy.data.meshes.new(object_name)
object = bpy.data.objects.new(object_name, mesh)

# Add edges. Using verts list to get start and end points for edges
for idx in range(len(verts)-1):
    edge = idx, idx+1
    edges.append(edge)

# create mesh from python data
mesh.from_pydata(verts, edges, [])
mesh.update(calc_edges=True)

# set mesh location
bpy.context.collection.objects.link(object)
