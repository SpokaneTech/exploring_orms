# Generated by Django 5.0.7 on 2024-07-15 23:52

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Manufacturer',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.TextField()),
            ],
        ),
        migrations.CreateModel(
            name='Person',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.TextField()),
            ],
        ),
        migrations.CreateModel(
            name='Model',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.TextField()),
                ('manufacturer', models.ForeignKey(on_delete=django.db.models.deletion.RESTRICT, to='cars.manufacturer')),
            ],
        ),
        migrations.CreateModel(
            name='Part',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.TextField()),
                ('cost', models.DecimalField(decimal_places=2, max_digits=7)),
                ('models', models.ManyToManyField(to='cars.model')),
            ],
        ),
        migrations.CreateModel(
            name='Vehicle',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('vin', models.TextField(null=True)),
                ('model', models.ForeignKey(on_delete=django.db.models.deletion.RESTRICT, to='cars.model')),
                ('person', models.OneToOneField(null=True, on_delete=django.db.models.deletion.RESTRICT, to='cars.person')),
            ],
        ),
    ]
