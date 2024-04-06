# Generated by Django 5.0.4 on 2024-04-06 16:44

import django.db.models.deletion
import uuid
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('currency', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='CurrencyRecord',
            fields=[
                ('id', models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True, serialize=False)),
                ('value', models.FloatField()),
                ('record_date', models.DateField()),
                ('created_at', models.DateField(auto_now_add=True)),
                ('updated_at', models.DateField(auto_now=True)),
                ('currency', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='currency.currency')),
            ],
        ),
    ]
