from django.urls import path
from . import views
from rest_framework import routers, serializers, viewsets, filters
from pages.models import Gist
# Serializers define the API representation.



class GistSerializer (serializers.ModelSerializer):
    class Meta:
        model = Gist
        fields = ['title','description','text']


class GistViewSet(viewsets.ModelViewSet):
    queryset = Gist.objects.all()
    serializer_class = GistSerializer
    http_method_names = ['get']
    search_fields = ['title','description','text']
    filter_backends = (filters.SearchFilter,)


router = routers.DefaultRouter()
router.register(r'gists', GistViewSet)

urlpatterns = [
    path("", views.home, name='home'),
    path("apps/<int:num>", views.app, name='apps'),
]

urlpatterns += router.urls