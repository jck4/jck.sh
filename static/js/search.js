

$(document).ready(function () {


    $("input").keyup(function () {
        var searchBoxContents = document.getElementById('search');
        var searchVal = searchBoxContents.value
        var params = { search: searchVal };

        var url = new URL('https://www.jck.sh/gists/')
        url.search = new URLSearchParams(params).toString();

        var gistCards = jQuery('<div>', {
            id: 'cards',
            class: 'gistCards',
        })


        fetch(url)
            .then(response => response.json())
            .then(gists => {

                var gistCard = jQuery('<div>', {
                    id: 'card',
                    class: 'gistCard',
                })

                if (gists.length == 0) {
                    console.log("yeet")
                    $(".content").html("No Gists Found");
                }

                for (var i in gists){
                    const card = `
                    <div class="card" id="card">
                        <h2 class="title">${gists[i].title}</h2>
                        <div class="description">${gists[i].description}</div>
                        <div class="text">${gists[i].text}</div>
                    </div>`;

                    gistCards.append(card)
                }
                $(".content").html(gistCards);
            })
            .catch(err => console.log(err))

    });
});

