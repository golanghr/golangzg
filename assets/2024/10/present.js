
let currentBottom = -75
let currentOpacity = 0
moveGopherToPlace = () => {
    currentBottom = -85
    currentOpacity = 0
    let podNameElement = document.getElementById('gopher-talk-id')
    let presenterElement = document.getElementById('presenter-talk-id')
    if (podNameElement === null) {
        setTimeout(moveGopherToPlace, 10);
        return;
    }
    presenterElement.style.opacity = 0
    podNameElement.style.bottom = "-25svh"
    let targetBottom = -2;

    let animationInterval = setInterval(() => {
      currentBottom += 0.1;
      podNameElement.style.bottom = `${currentBottom}svh`;
      if (currentBottom >= targetBottom) {
          clearInterval(animationInterval);
          showPresenter()
      }
    }, 3);

}
showPresenter = () => {
    let presenterElement = document.getElementById('presenter-talk-id')
    let animationInterval2 = setInterval(() => {
        // Increment currentBottom by a small amount (adjust as needed)
        currentOpacity += 0.01;
        // Update the bottom property
        presenterElement.style.opacity = `${currentOpacity}`;
        // Check if the target bottom is reached
        if (currentOpacity >= 1) {
            clearInterval(animationInterval2);
        }
      }, 10);
}

(function() {
    var beforePrint = function() {
        let podNameElement = document.getElementById('gopher-talk-id')
        podNameElement.style.bottom = `0svh`;
    };
    var afterPrint = function() {
        let podNameElement = document.getElementById('gopher-talk-id')
        podNameElement.style.bottom = `-2svh`;
    };

    if (window.matchMedia) {
        var mediaQueryList = window.matchMedia('print');
        mediaQueryList.addListener(function(mql) {
            if (mql.matches) {
                beforePrint();
            } else {
                afterPrint();
            }
        });
    }

    window.onbeforeprint = beforePrint;
    window.onafterprint = afterPrint;
}());
