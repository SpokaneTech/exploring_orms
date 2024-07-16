from django.core.management.base import BaseCommand, CommandError
from cars import models


class Command(BaseCommand):
    def add_arguments(self, parser):
        parser.add_argument("--vehicle_id", type=int, required=True)
        parser.add_argument("--name", type=str, required=True)

    def handle(self, **options):
        vehicle_id: int = options["vehicle_id"]
        person_name: str = options["name"]

        vehicle = (
            models.Vehicle.objects.select_related("model", "model__manufacturer")
            .filter(pk=vehicle_id)
            .first()
        )
        if not vehicle:
            raise CommandError("No vehicle found with ID {0}".format(vehicle_id))

        person, _ = models.Person.objects.filter(name=person_name).get_or_create(
            {"name": person_name}
        )
        vehicle.person_id = person.pk
        vehicle.save()

        print(
            "Added a {0} as the owner of ({1}) {2} {3}".format(
                person.name,
                vehicle.pk,
                vehicle.model.name,
                vehicle.model.manufacturer.name,
            )
        )
