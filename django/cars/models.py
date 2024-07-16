from __future__ import annotations
from django.db import models


class Manufacturer(models.Model):
    """A vehicle manufacturer, like Chevrolet"""

    name = models.TextField()

    vehicles: models.ManyToManyRel


class Model(models.Model):
    """A vehicle model, like a Chevrolet Silverado"""

    name = models.TextField()
    manufacturer = models.ForeignKey(Manufacturer, on_delete=models.RESTRICT)

    parts: models.ManyToManyRel


class Vehicle(models.Model):
    """An individual of a model, like Joe's Chevrolet Silverado"""

    vin = models.TextField(null=True)
    model = models.ForeignKey(Model, on_delete=models.RESTRICT)
    person = models.OneToOneField("cars.Person", on_delete=models.RESTRICT, null=True)

    person_id: int | None


class Part(models.Model):
    """A vehicle part for one or more models, like a muffler for all Chevrolet pickups"""

    name = models.TextField()
    cost = models.DecimalField(max_digits=7, decimal_places=2)
    models = models.ManyToManyField(Model)


class Person(models.Model):
    """A person, who may drive a vehicle"""

    name = models.TextField()

    vehicle: Vehicle | None
