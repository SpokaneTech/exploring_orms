from django.core.management.base import BaseCommand
from cars import models


class Command(BaseCommand):
    def handle(self, *args, **options):
        vehicles = models.Vehicle.objects.select_related("model", "model__manufacturer")
        if len(vehicles) == 0:
            print("No vehicles found in your garage")
        for vehicle in vehicles:
            print(
                "({0}) {1} {2}".format(
                    vehicle.pk,
                    vehicle.model.name,
                    vehicle.model.manufacturer.name,
                )
            )
