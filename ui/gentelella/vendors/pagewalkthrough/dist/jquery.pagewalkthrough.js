/***
 * Author: Erwin Yusrizal
 * UX by: Tai Nguyen
 * Contributors: James West <jwwest@gmail.com>
 *               JP Adams <jpadamsonline@gmail.com>
 *               James Warwood <james.duncan.1991@googlemail.com>
 *               Craig Roberts <craig0990@googlemail.com>
 * Created On: 27/02/2013
 * Version: 2.7.2
 * Features & Bugs: https://github.com/warby-/jquery-pagewalkthrough/issues
 ***/

;(function($, window, document, undefined) {
  'use strict';

  /**
   * GLOBAL VAR
   */
  var _globalWalkthrough = {},
    _elements = [],
    _activeWalkthrough,
    _activeId,
    _hasDefault = true,
    _counter = 0,
    _isCookieLoad,
    _firstTimeLoad = true,
    _onLoad = true,
    _index = 0,
    _isWalkthroughActive = false,
    $jpwOverlay = $('<div id="jpwOverlay"></div>'),
    $jpWalkthrough = $('<div id="jpWalkthrough"></div>'),
    $jpwTooltip = $('<div id="jpwTooltip"></div>');

  /**
   * PUBLIC METHOD
   */

  var methods = {
    isActive: function() {
      return !!_isWalkthroughActive;
    },

    index: function(value) {
      if (typeof value !== 'undefined') {
        _index = value;
      }

      return _index;
    },

    //init method
    init: function(options) {
      options = $.extend(true, {}, $.fn.pagewalkthrough.defaults, options);
      var that = this;

      if (!options.name) {
        throw new Error('Must provide a unique name for a tour');
      }

      // @todo what happens with multiple walkthroughs on the same element?
      this.first().data('jpw', options);

      options._element = this;

      return this.each(function(i) {
        options = options || {};
        options.elementID = options.name;

        _globalWalkthrough[options.name] = options;
        _elements.push(options.name);

        //check if onLoad and this is first time load
        if (options.onLoad) {
          _counter++;
        }

        //get first onload = true
        if (_counter === 1 && _onLoad) {
          _activeId = options.name;
          _activeWalkthrough = _globalWalkthrough[_activeId];
          _onLoad = false;
        }

        // set the activeWalkthrough if onLoad is false for all walkthroughs
        if ((i + 1 === that.length && _counter === 0)) {
          _activeId = options.name;
          _activeWalkthrough = _globalWalkthrough[_elements[0]];
          _hasDefault = false;
        }
      });
    },

    renderOverlay: function() {

      // if each walkthrough has onLoad: true, log warning message
      if (_counter > 1) {
        debug('Warning: Only 1st walkthrough will be shown onLoad as default');
      }

      //get cookie load
      _isCookieLoad = getCookie('_walkthrough-' + _activeId);

      //check if first time walkthrough
      if (typeof _isCookieLoad === 'undefined') {
        _isWalkthroughActive = true;

        if (!(onEnter())) {
          return;
        }

        showStep();
        showButton('jpwClose', 'body');

        setTimeout(function() {
          //call onAfterShow callback
          if (isFirstStep() && _firstTimeLoad) {
            if (!onAfterShow()) {
              return;
            }
          }
        }, 100);
      } else {
        onCookieLoad(_globalWalkthrough);
      }
    },

    restart: function(e) {
      if (isFirstStep()) {
        return;
      }

      _index = 0;
      if (!(onRestart(e)) || !(onEnter(e))) {
        return;
      }

      showStep();
    },

    close: function() {
      var options = _activeWalkthrough;

      onLeave(true);

      if (typeof options.onClose === 'function') {
        options.onClose.call(this);
      }

      _index = 0;
      _firstTimeLoad = true;

      _isWalkthroughActive = false;

      //set cookie to false
      setCookie('_walkthrough-' + _activeId, 0, 365);
      _isCookieLoad = getCookie('_walkthrough-' + _activeId);

      $jpwOverlay.fadeOut('slow', function() {
        $(this).remove();
      });

      $jpWalkthrough.fadeOut('slow', function() {
        $(this).html('').remove();
      });

      $('#jpwClose').fadeOut('slow', function() {
        $(this).remove();
      });

    },

    show: function(name, e) {
      // If no name, then first argument is event
      e = name == null ? name : e;

      name = name || this.first().data('jpw').name;

      _activeWalkthrough = _globalWalkthrough[this.first().data('jpw').name];

      if ((name === _activeId && _isWalkthroughActive) || !(onEnter(e))) {
        return;
      }

      _isWalkthroughActive = true;
      _firstTimeLoad = true;
      if (!(onBeforeShow())) {
        return;
      }

      showStep();
      showButton('jpwClose', 'body');

      //call onAfterShow callback
      if ((isFirstStep() && _firstTimeLoad) && !onAfterShow()) {
        return;
      }
    },

    next: function(e) {
      _firstTimeLoad = false;
      if (isLastStep() || !onLeave(e)) {
        return;
      }

      _index = parseInt(_index, 10) + 1;
      if (!onEnter(e)) {
          methods.next();
      }
      showStep('next');
    },

    prev: function(e) {
      if (isFirstStep() || !onLeave(e)) {
        return;
      }

      _index = parseInt(_index, 10) - 1;
      if (!onEnter(e)) {
        methods.prev();
      }
      showStep('prev');
    },

    getOptions: function(activeWalkthrough) {
      var _wtObj;

      //get only current active walkthrough
      if (activeWalkthrough) {
        if (!_isWalkthroughActive) {
          _wtObj = false;
        } else {
          _wtObj = _activeWalkthrough;
        }
        //get all walkthrough
      } else {
        _wtObj = [];
        for (var wt in _globalWalkthrough) {
          _wtObj.push(_globalWalkthrough[wt]);
        }
      }

      return _wtObj;
    },

    refresh: function() {
        // Stricly speaking, a skipDirection should never
        // be needed, but I'd rather provide one at this point
        // than watch it explode...
        showStep('next');
    }
  }; //end public method

  /* Pre-build walkthrough step function.  Handles the scrolling to the target
   * element.
   */
  function showStep(skipDirection) {
    var options = _activeWalkthrough,
      step = options.steps[_index],
      targetElement = options._element.find(step.wrapper),
      scrollTarget = getScrollParent(targetElement),
      maxScroll, scrollTo;

      if (step.popup.type !== 'modal' && !targetElement.length) {
        if (step.popup.fallback === 'skip' ||
            typeof step.popup.fallback === 'undefined') {
          methods[skipDirection]();
          return;
        }

        step.popup.type = step.popup.fallback;
      }

      // For modals, scroll to the top.  For tooltips, try and center the target
      // (wrapper) element in the screen
      maxScroll = scrollTarget[0].scrollHeight - scrollTarget.outerHeight();

      if (step.autoScroll !== false) {
        if (step.popup.type === 'modal') {
          scrollTo = 0;
        } else {
          scrollTo = Math.floor(
            targetElement.offset().top - ($(window).height() / 2) +
                scrollTarget.scrollTop()
          );
        }
      } else {
        scrollTo = scrollTarget.scrollTop();
      }

    // @TODO: simplify this logic
    //
    // Conditions for scrolling:
    //   1.  new scroll value is not equal to current scroll value
    //    AND
    //     a.  new scroll value is less than the max scroll, and we are
    //         currently at the max scroll value
    //      OR
    //     b.  new scroll value is less than or equal to 0, and the current
    //         scroll is greater than 0
    //      OR
    //     c.  new scroll value is greater than 0
    if (scrollTarget.scrollTop() !== scrollTo &&
      (
        (scrollTarget.scrollTop() === maxScroll && scrollTo < maxScroll) ||
        (scrollTo <= 0 && scrollTarget.scrollTop() > 0) ||
        (scrollTo > 0)
      )) {

      // Stylistic concerns - fill overlay hole and hide tooltip whilst
      // scrolling
      $jpWalkthrough.addClass('jpw-scrolling');
      $jpwTooltip.fadeOut('fast');

      scrollTarget.clearQueue().animate({
        scrollTop: scrollTo
      }, options.steps[_index].scrollSpeed, buildWalkthrough);

    } else {

      // Not scrolling, so jump directly to building the walkthrough
      buildWalkthrough();
    }
  }

  function buildWalkthrough() {
    $jpWalkthrough.removeClass('jpw-scrolling');
    var options = _activeWalkthrough,
      step = options.steps[_index],
      targetElement,
      scrollParent,
      maxHeight;

    // Extend step options with defaults
    options.steps[_index] = $.extend(
        true, {}, $.fn.pagewalkthrough.defaults.steps[0], step
    );

    targetElement = options._element.find(step.wrapper);
    scrollParent = getScrollParent(targetElement);

    $jpwOverlay.show();

    if (step.popup.type !== 'modal' && step.popup.type !== 'nohighlight') {

      $jpWalkthrough.html('');

      //check if wrapper is not empty or undefined
      if (step.wrapper === '' || typeof step.wrapper === 'undefined') {
        // @TODO should we skip here?
        debug('Your walkthrough position is: "' + step.popup.type +
            '" but wrapper is empty or undefined. Please check your "' +
            _activeId + '" wrapper parameter.'
        );
        return;
      }

      maxHeight = scrollParent.outerHeight() - targetElement.offset().top +
        scrollParent.offset().top + scrollParent.scrollTop();

      // If max height is negative, means we have plenty room so targetElement
      // height
      maxHeight = maxHeight <= 0 ? targetElement.outerHeight() : maxHeight;

      // @todo make it so we don't have to destroy and recreate this element for
      // each step
      $jpwOverlay.appendTo($jpWalkthrough);

      // Overlay hole
      $('<div>')
        .addClass('overlay-hole')
        .height(Math.min(maxHeight, targetElement.outerHeight())+10)
        .width(targetElement.outerWidth())
        .css({
             // Recommended to be at least twice the inset box-shadow spread
            padding: '20px',
            position: 'absolute',
            top: targetElement.offset().top, // top/left minus padding
            left: targetElement.offset().left,
            'z-index': 999998,
            'box-shadow': '0 0 1px 10000px rgba(0, 0, 0, 0.6)'
        })
        .append(
            $('<div>')
                .css({
                    position: 'absolute',
                    top: 0,
                    bottom: 0,
                    left: 0,
                    right: 0
                })
        )
        .appendTo($jpWalkthrough);

      if ($('#jpWalkthrough').length) {
        $('#jpWalkthrough').remove();
      }

      $jpWalkthrough.appendTo('body').show();
      $jpwTooltip.show();

      showTooltip();

    } else if (step.popup.type === 'modal') {

      if ($('#jpWalkthrough').length) {
        $('#jpWalkthrough').remove();
      }

      showModal();

    } else {
      if ($('#jpWalkthrough').length) {
        $('#jpWalkthrough').remove();
      }
    }

    showButton('jpwPrevious');
    showButton('jpwNext');
    showButton('jpwFinish');
  }

  /*
   * SHOW MODAL
   */

  function showModal() {
    var options = _activeWalkthrough,
      step = options.steps[_index];

    $jpwOverlay.appendTo('body').show().removeClass('transparent');

    var textRotation = setRotation(parseInt(step.popup.contentRotation, 10));

    $jpwTooltip.css({
      'position': 'absolute',
      'left': '50%',
      'top': 'calc('+$(document).scrollTop()+'px + 25%)',
      'margin-left': -(parseInt(step.popup.width, 10) + 60) / 2 + 'px',
      'z-index': '999999'
    });

    var tooltipSlide = $('<div id="tooltipTop">' +
      '<div id="topLeft"></div>' +
      '<div id="topRight"></div>' +
      '</div>' +

    '<div id="tooltipInner">' +
      '</div>' +

    '<div id="tooltipBottom">' +
      '<div id="bottomLeft"></div>' +
      '<div id="bottomRight"></div>' +
      '</div>');

    $jpWalkthrough.html('');
    $jpwTooltip.html('').append(tooltipSlide)
      .wrapInner($('<div />', {
        id: 'tooltipWrapper',
        style: 'width:' + cleanValue(parseInt(step.popup.width, 10) + 30 + 120)
      }))
      .append('<div id="bottom-scratch"></div>')
      .appendTo($jpWalkthrough);

    $jpWalkthrough.appendTo('body');
    $jpwTooltip.show();

    $('#tooltipWrapper').css(textRotation);

    $('#tooltipInner').append(getContent(step)).show();

    $jpWalkthrough.show();
  }


  /*
   * SHOW TOOLTIP
   */

  function showTooltip() {
    var opt = _activeWalkthrough,
      step = opt.steps[_index];

    var top, left, arrowTop, arrowLeft,
      overlayHoleWidth = $('#jpWalkthrough .overlay-hole').outerWidth(),
      overlayHoleHeight = $('#jpWalkthrough .overlay-hole').outerHeight(),
      overlayHoleTop = $('#jpWalkthrough .overlay-hole').offset().top,
      overlayHoleLeft = $('#jpWalkthrough .overlay-hole').offset().left,
      arrow = 30;

    var textRotation = (typeof step.popup.contentRotation === 'undefined' ||
        parseInt(step.popup.contentRotation, 10) === 0) ? clearRotation() :
        setRotation(parseInt(step.popup.contentRotation, 10));


    // Remove overlay background to prevent double-transparency
    if ($('#jpwOverlay').length) {
      $('#jpwOverlay').addClass('transparent');
    }

    var tooltipSlide = $('<div id="tooltipTop">' +
      '<div id="topLeft"></div>' +
      '<div id="topRight"></div>' +
      '</div>' +

    '<div id="tooltipInner">' +
      '</div>' +

    '<div id="tooltipBottom">' +
      '<div id="bottomLeft"></div>' +
      '<div id="bottomRight"></div>' +
      '</div>');

    $jpwTooltip.html('').css({
      'margin-left': '0',
      'margin-top': '0',
      'position': 'absolute',
      'z-index': '999999'
    })
      .append(tooltipSlide)
      .wrapInner($('<div />', {
          id: 'tooltipWrapper',
          style: 'width:' + cleanValue(parseInt(step.popup.width, 10) + 30 + 120)
      }))
      .appendTo($jpWalkthrough);

    $jpWalkthrough.appendTo('body').show();

    $('#tooltipWrapper').css(textRotation);

    $('#tooltipInner').append(getContent(step)).show();

    $jpwTooltip.append(
        '<span class="' + step.popup.position + '">&nbsp;</span>'
    );

    switch (step.popup.position) {

      case 'top':
        top = overlayHoleTop - ($jpwTooltip.height() + (arrow / 2)) +
            parseInt(step.popup.offsetVertical, 10) - 86;
        left = (overlayHoleLeft + (overlayHoleWidth / 2)) -
          ($jpwTooltip.width() / 2) - 5 +
          parseInt(step.popup.offsetHorizontal, 10);
        arrowLeft = ($jpwTooltip.width() / 2) - arrow +
            parseInt(step.popup.offsetArrowHorizontal, 10);
        arrowTop = (step.popup.offsetArrowVertical) ?
            parseInt(step.popup.offsetArrowVertical, 10) :
            '';
        break;
      case 'right':
        top = overlayHoleTop - (arrow / 2) +
            parseInt(step.popup.offsetVertical, 10);
        left = overlayHoleLeft + overlayHoleWidth + (arrow / 2) +
            parseInt(step.popup.offsetHorizontal, 10) + 105;
        arrowTop = arrow + parseInt(step.popup.offsetArrowVertical, 10);
        arrowLeft = (step.popup.offsetArrowHorizontal) ?
            parseInt(step.popup.offsetArrowHorizontal, 10) :
            '';
        break;
      case 'bottom':
        top = overlayHoleTop + overlayHoleHeight +
          parseInt(step.popup.offsetVertical, 10) + 86;
        left = (overlayHoleLeft + (overlayHoleWidth / 2)) -
          ($jpwTooltip.width() / 2) - 5 +
          parseInt(step.popup.offsetHorizontal, 10);

        arrowLeft = (($jpwTooltip.width() / 2) - arrow) +
            parseInt(step.popup.offsetArrowHorizontal, 10);
        arrowTop = (step.popup.offsetArrowVertical) ?
            parseInt(step.popup.offsetArrowVertical, 10) :
            '';
        break;
      case 'left':
        top = overlayHoleTop - (arrow / 2) +
            parseInt(step.popup.offsetVertical, 10);
        left = overlayHoleLeft - $jpwTooltip.width() - (arrow) +
            parseInt(step.popup.offsetHorizontal, 10) - 105;
        arrowTop = arrow + parseInt(step.popup.offsetArrowVertical, 10);
        arrowLeft = (step.popup.offsetArrowVertical) ?
            parseInt(step.popup.offsetArrowHorizontal, 10) :
            '';
        break;
    }

    $('#jpwTooltip span.' + step.popup.position).css({
      'top': cleanValue(arrowTop),
      'left': cleanValue(arrowLeft)
    });

    $jpwTooltip.css({
      'top': cleanValue(top),
      'left': cleanValue(left)
    });
    $jpWalkthrough.show();
  }

  /* Get the content for a step.  First attempts to treat step.popup.content
   * as a selector.  If this fails, or returns an empty result set, it falls
   * back to return the value of step.popup.content.
   *
   * This allows both selectors and literal content to be provided in the
   * content option.
   *
   * @param {Object} step  The step data to return the content for
   */
  function getContent(step) {
    var option = step.popup.content,
      content;

    try {
      content = $('body').find(option).html();
    } catch(e) {
    }

    return content || option;
  }

  /* Render a control button outside the #tooltipInner element.
   *
   * @param {String} id               The button identifier within the
   *                                  options.buttons hash (e.g. 'jpwNext')
   * @param {jQuery|String} appendTo  (Optional) The element or selector to
   *                                  append the button to.  Defaults to
   *                                  #tooltipWrapper
   */
  function showButton(id, appendTo) {
    if ($('#' + id).length) {
      return;
    }

    var btn = _activeWalkthrough.buttons[id],
      $a;

    // Check that button is defined
    if (!btn) {
      return;
    }

    // Check that button should be shown
    if ((typeof btn.show === 'function' && !btn.show()) || !btn.show) {
      return;
    }

    $a = $('<a />', {
      id: id,
      html: btn.i18n
    });

    // Append button
    if (appendTo) {
      $(appendTo).append($a);
    } else {
      $('#tooltipWrapper').after($a);
    }
  }

  /**
    /* CALLBACK
    /*/

  //callback for onLoadHidden cookie

  function onCookieLoad(options) {
    /*jshint validthis: true */
    for (var i = 0; i < _elements.length; i++) {
      if (typeof(options[_elements[i]].onCookieLoad) === 'function') {
        options[_elements[i]].onCookieLoad.call(this);
      }
    }

    return false;
  }

  function onLeave(e) {
    /*jshint validthis: true */
    var options = _activeWalkthrough;

    if (typeof options.steps[_index].onLeave === 'function') {
      if (options.steps[_index].onLeave.call(this, e, _index) === false) {
        return false;
      }
    }

    return true;

  }

  //callback for onEnter step

  function onEnter(e) {
    /*jshint validthis: true */
    var options = _activeWalkthrough;

    if (typeof options.steps[_index].onEnter === 'function') {
      return options.steps[_index].onEnter.call(this, e, _index) !== false;
    }

    return true;
  }

  //callback for onRestart help

  function onRestart(e) {
    /*jshint validthis: true */
    var options = _activeWalkthrough;
    //set help mode to true
    _isWalkthroughActive = true;
    methods.restart(e);

    if (typeof options.onRestart === 'function') {
      if (options.onRestart.call(this) === false) {
        return false;
      }
    }

    return true;
  }

  //callback for before all first walkthrough element loaded

  function onBeforeShow() {
    /*jshint validthis: true */
    var options = _activeWalkthrough || {};
    _index = 0;

    if (typeof(options.onBeforeShow) === 'function') {
      if (options.onBeforeShow.call(this) === false) {
        return false;
      }
    }

    return true;
  }

  //callback for after all first walkthrough element loaded

  function onAfterShow() {
    /*jshint validthis: true */
    var options = _activeWalkthrough;
    _index = 0;

    if (typeof(options.onAfterShow) === 'function') {
      if (options.onAfterShow.call(this) === false) {
        return false;
      }
    }

    return true;
  }



  /**
   * HELPERS
   */
  function debug(message) {
    if (window.console && window.console.log) {
      window.console.log(message);
    }
  }

  function clearRotation() {
    var rotationStyle = {
      '-webkit-transform': 'none', //safari
      '-moz-transform': 'none', //firefox
      '-o-transform': 'none', //opera
      'filter': 'none', //IE7
      '-ms-transform': 'none' //IE8+
    };

    return rotationStyle;
  }

  function setRotation(angle) {

    //for IE7 & IE8
    var M11, M12, M21, M22, deg2rad, rad;

    //degree to radian
    deg2rad = Math.PI * 2 / 360;
    rad = angle * deg2rad;

    M11 = Math.cos(rad);
    M12 = Math.sin(rad);
    M21 = Math.sin(rad);
    M22 = Math.cos(rad);

    var rotationStyle = {
      '-webkit-transform': 'rotate(' + parseInt(angle, 10) + 'deg)', //safari
      '-moz-transform': 'rotate(' + parseInt(angle, 10) + 'deg)', //firefox
      '-o-transform': 'rotate(' + parseInt(angle, 10) + 'deg)', //opera
      '-ms-transform': 'rotate(' + parseInt(angle, 10) + 'deg)' //IE9+
    };

    return rotationStyle;

  }

  function cleanValue(value) {
    if (typeof value === 'string') {
      if (value.toLowerCase().indexOf('px') === -1) {
        return value + 'px';
      } else {
        return value;
      }
    } else {
      return value + 'px';
    }
  }

  function setCookie(cName, value, exdays) {
    var exdate = new Date();
    exdate.setDate(exdate.getDate() + exdays);
    var cValue = encodeURIComponent(value) + ((exdays == null) ? '' :
        '; expires=' + exdate.toUTCString());
    document.cookie = [cName, '=', cValue].join('');
  }

  function getCookie(cName) {
    var i, x, y, ARRcookies = document.cookie.split(';');
    for (i = 0; i < ARRcookies.length; i++) {
      x = ARRcookies[i].substr(0, ARRcookies[i].indexOf('='));
      y = ARRcookies[i].substr(ARRcookies[i].indexOf('=') + 1);
      x = x.replace(/^\s+|\s+$/g, '');
      if (x === cName) {
        return decodeURIComponent(y);
      }
    }
  }

  /* Returns true if the current step is the last step in the walkthrough.
   *
   * @return {Boolean} true if user is currently viewing last step; false
   *                   otherwise
   */
  function isLastStep() {
    return _index === (_activeWalkthrough.steps.length - 1);
  }

  /* Returns true if the current step is the first step in the walkthrough.
   *
   * @return {Boolean} true if user is currently viewing first step; false
   *                   otherwise
   */
  function isFirstStep() {
      return _index === 0;
  }

  /* Get the first scrollable parent of the specified element.
   * Adapted from jQueryUI's [scrollParent](api.jqueryui.com/scrollParent/).
   *
   * @param {jQuery} element  The element to find the scrollable parent of
   *
   * @return {jQuery} The first scrollable parent, or an empty jQuery object if
   *                  either of the following is true:
   *                    1. `element`'s position is 'fixed'
   *                    2. `element`'s position is 'absolute', and the parent's
   *                        is 'static'
   */
  function getScrollParent(element) {
    if (!(element instanceof $)) {
      element = $(element);
    }

    element = element.first();

    var position = element.css('position'),
      excludeStaticParent = position === 'absolute',
      scrollParent = element.parents().filter(function() {
        var parent = $(this);
        if (excludeStaticParent && parent.css('position') === 'static') {
          return false;
        }

        return (/(auto|scroll)/).test(
          parent.css('overflow') + parent.css('overflow-y') +
          parent.css('overflow-x')
        );
      }).eq(0);

    return position === 'fixed' ? $() : !scrollParent.length ?
      $('html, body') : scrollParent;
  }

  /**
   * BUTTON CLOSE CLICK
   */

  /* Close and finish tour buttons clicks */
  $(document).on('click', '#jpwClose, #jpwFinish', methods.close);

  /* Next button clicks
   */
  $(document).on('click', '#jpwNext', function() {
    $.pagewalkthrough('next');
  });

  /* Previous button clicks
   */
  $(document).on('click', '#jpwPrevious', function() {
    $.pagewalkthrough('prev');
  });

  $(document).on(
    'click',
    '#jpwOverlay, #jpwTooltip',
    function(ev) {
        ev.stopPropagation();
        ev.stopImmediatePropagation();
    }
  );

  /**
   * WINDOW RESIZE RERENDERER
   */
  
  $(window).resize(function() {
    if (_isWalkthroughActive) {
      $.pagewalkthrough('refresh');
    }
  });

  /**
   * DRAG & DROP
   */

  /**
   * MAIN PLUGIN
   */
  $.pagewalkthrough = $.fn.pagewalkthrough = function(method) {

    if (methods[method]) {

      return methods[method].apply(this, [].slice.call(arguments, 1));

    } else if (typeof method === 'object' || !method) {

      methods.init.apply(this, arguments);

      // render the overlay on it has a default walkthrough set to show onload
      if (_hasDefault && _counter < 2) {
        setTimeout(function() {
          methods.renderOverlay();
        }, 500);
      }

    } else {

      $.error('Method ' + method + ' does not exist on jQuery.pagewalkthrough');

    }

  };

  /* #### <a name="default-options">Options</a>
   *
   * Default options for each walkthrough.
   * User options extend these defaults.
   */
  $.fn.pagewalkthrough.defaults = {
    /* Array of steps to show
     */
    steps: [
      {
        // jQuery selector for the element to highlight for this step
        wrapper: '',
        // ##### <a name="popup-options">Popup options</a>
        popup: {
          // Selector for the element which contains the content, or the literal
          // content
          content: '',
          // Popup type - either modal, tooltip or nohighlight.
          // See [Popup Types](/pages/popup-types.html)
          type: 'modal',
          // Position for tooltip and nohighlight style popups - either top,
          // left, right or bottom
          position: 'top',
          // Horizontal offset for the walkthrough
          offsetHorizontal: 0,
          // Vertical offset for the walkthrough
          offsetVertical: 0,
          // Horizontal offset for the arrow
          offsetArrowHorizontal: 0,
          // Vertical offset for the arrow
          offsetArrowVertical: 0,
          // Default width for each popup
          width: '320',
          // Amount in degrees to rotate the content by
          contentRotation: 0,
          // If set to 'skip', skips tooltip/nohighlight types when the wrapper
          // does not exist. If set to anything else, uses the value as a fallback
          // popup type (e.g. 'modal' will fallback to a popup type of 'modal').
          fallback: 'skip'
        },
        // Automatically scroll to the content for the step
        autoScroll: true,
        // Speed to use when scrolling to elements
        scrollSpeed: 1000,
        // Callback when entering the step
        onEnter: $.noop,
        /* Callback when leaving the step.  Called with `true` if the user is
         * skipping the rest of the tour (gh #66)
         */
        onLeave: $.noop
      }
    ],
    // **(Required)** Walkthrough name.  Should be a unique name to identify the
    // walkthrough, as it will
    // be used in the cookie name
    name: null,
    // Automatically show the walkthrough when the page is loaded.  If multiple
    // walkthroughs set this to true, only the first walkthrough is shown
    // automatically
    onLoad: true,
    // Callback to be executed before the walkthrough is shown
    onBeforeShow: $.noop,
    // Callback executed after the walkthrough is shown
    onAfterShow: $.noop,
    // Callback executed in the event that 'restart' is triggered
    onRestart: $.noop,
    // Callback executed when the walkthrough is closed.  The walkthrough can be
    // closed by the user clicking the close button in the top right, or
    // clicking the finish button on the last step
    onClose: $.noop,
    // Callback executed when cookie has been set after a walkthrough has been
    // closed
    onCookieLoad: $.noop,
    /* ##### <a name="controls-options">Walkthrough controls</a>
     *
     * Hash of buttons to show.  Object keys are used as the button element's ID
     */
    buttons: {
      // ID of the button
      jpwClose: {
        // Translation string for the button
        i18n: '关闭',
        // Whether or not to show the button.  Can be a boolean value, or a
        // function which returns a boolean value
        show: true
      },
      jpwNext: {
        i18n: '下一条 &rarr;',
        // Function which resolves to a boolean
        show: function() {
          return !isLastStep();
        }
      },
      jpwPrevious: {
        i18n: '&larr; 上一条',
        show: function() {
          return !isFirstStep();
        }
      },
      jpwFinish: {
        i18n: '完成 &#10004;',
        show: function() {
          return isLastStep();
        }
      }
    }
  };
}(jQuery, window, document));

