from django.shortcuts import render
from django.http import Http404
from django.template.loader import render_to_string
# Create your views here.

templatesDirs = ['index.html','generator.html','terminal.html','gists.html']

def home(request):
    return render(request,'index.html')

def app(request, num):
    
    if templatesDirs[num]:
        return render(request, templatesDirs[num])
    else:
        raise Http404("No such app")
