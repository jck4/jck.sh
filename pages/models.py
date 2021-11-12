from django.db import models

# Create your models here.
class Gist(models.Model):
    title = models.CharField(max_length=100)
    description = models.CharField(max_length=250)
    text = models.TextField(max_length=2048)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)