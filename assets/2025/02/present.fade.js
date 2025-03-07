document.querySelectorAll('.slide').forEach(function(slide) {
  slide.style.position = 'absolute';
  slide.style.top = 0;
  slide.style.left = '100%';
});

document.querySelectorAll('.slide.hidden').forEach(function(slide) {
  slide.classList.remove('hidden');
  slide.style.opacity = 0;
});

slideAnimationTime = 800;
easingType = 'easeInOutCubic';
easingType = 'linear';
oldPage = -1

updateSlideVisibility = function (page) {
  if (oldPage == -1) {
    oldPage = page;
    getSlideElements().forEach(function(slide) {
      if (parseInt(slide.id.slice(6)) == page) {
        slide.style.opacity = 1;
        slide.style.left = '0';
        slide.classList.add('visible')
      }
    });
    return;
  }
  getSlideElements().forEach(function(slide) {
    if (parseInt(slide.id.slice(6)) != page && slide.classList.contains('visible')) {
      //slide.classList.add('hidden');
      slide.classList.remove('visible');
      anime({
        targets: document.querySelector('#slide-'+parseInt(slide.id.slice(6))),
        opacity: 0,
        duration: slideAnimationTime,
        easing: easingType,
        complete: function(anim) {
          slide.style.left = '100%';
        }
      });
    }
  });
  getSlideElements().forEach(function(slide) {
    if (parseInt(slide.id.slice(6)) == page) {
      slide.style.left = '0';
      slide.classList.remove('hidden');
      slide.classList.add('visible');
      anime({
        targets: document.querySelector('#slide-'+page),
        opacity: 1,
        duration: slideAnimationTime,
        easing: easingType
      });
    }
  });
}
