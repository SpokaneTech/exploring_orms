from django.core.management.base import BaseCommand
from cars import models


class Command(BaseCommand):
    def add_arguments(self, parser):
        parser.add_argument("--manufacturer", type=str, required=True)
        parser.add_argument("--model", type=str, required=True)
        parser.add_argument("--vin", type=str, required=False)

    def handle(self, **options):
        manufacturer_name: str = options["manufacturer"]
        model_name: str = options["model"]
        vin: str | None = options["vin"]

        manufacturer, _ = models.Manufacturer.objects.get_or_create(
            {"name": manufacturer_name}
        )

        model, _ = models.Model.objects.get_or_create(
            {
                "name": model_name,
                "manufacturer": manufacturer,
            }
        )

        vehicle = models.Vehicle.objects.create(model=model, vin=vin)
        print(
            "Added a {0} {1} to your garage with the VIN {2}".format(
                manufacturer.name,
                model.name,
                vehicle.vin,
            )
        )
