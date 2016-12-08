# jQuery Page Walkthrough
[![GitHub version](https://badge.fury.io/gh/jwarby%2Fjquery-pagewalkthrough.svg)](http://badge.fury.io/gh/jwarby%2Fjquery-pagewalkthrough)
[![Build Status](https://travis-ci.org/jwarby/jquery-pagewalkthrough.png?branch=dev)](https://travis-ci.org/jwarby/jquery-pagewalkthrough)

Forked from [jpadamsonline/jquerypagewalkthrough.github.com](https://github.com/jpadamsonline/jquerypagewalkthrough.github.com)

Page Walkthrough is a flexible system for designing interactive, multimedia, educational walkthroughs.

**Note:** currently under heavy active development, and is likely to change **alot** at this moment in time.  Check out where we're at by going to
the [issues](https://github.com/jwarby/jquery-pagewalkthrough/issues) page.

## Screenshots

### Modal-style tour step
![Modal-style step](assets/screenshot_modal.png 'Modal-style step')

### Tooltip-style tour step with highlighted content
![Tooltip-style step](assets/screenshot_tooltip.png 'Tooltip-style step')

## Demo site

The demo site is located [here](http://jwarby.github.io/jquery-pagewalkthrough/).

## Installing

1. Download the release you want from the [releases page](https://github.com/jwarby/jquery-pagewalkthrough/releases), or download the
[latest code](https://github.com/jwarby/jquery-pagewalkthrough/archive/master.zip)(may not be stable).
2. Extract the files from the `dist/` folder into your project
3. Include the stylesheets and JS (**note: include jQuery first**):

```html
<!-- CSS -->
<link type="text/css" rel="stylesheet" href="path/to/extracted/files/css/jquery.pagewalkthrough.css" />

<!-- jQuery -->
<script type="text/javascript" src="path/to/jquery/jquery-<jquery_version>.js"></script>
<!-- Page walkthrough plugin -->
<script type="text/javascript" src="path/to/extracted/files/jquery.pagewalkthrough.js"></script>
```

**Note: there are minified versions of the JS and CSS files available for production, with the suffixes `.min.js` and `.min.css` respectively.**

## Usage

@TODO

## jQuery Page Walkthrough Default Options

**Note:** as of version 1.4, you **must** specify a tour name in the options.

**Note:** as of version 2.1, the `popup.content` option for each step can now be the literal content for the step, or a selector as before.  The rule
for how the plugin decides which to treat it as is simple: if the string is a valid selector, and the selector matches an element that is already present
in the DOM, the matched element's content is displayed; if it is an invalid selector, or a selector which returns no matches, the literal value is displayed.

```javascript
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
      i18n: 'Click here to close',
      // Whether or not to show the button.  Can be a boolean value, or a
      // function which returns a boolean value
      show: true
    },
    jpwNext: {
      i18n: 'Next &rarr;',
      // Function which resolves to a boolean
      show: function() {
        return !isLastStep();
      }
    },
    jpwPrevious: {
      i18n: '&larr; Previous',
      show: function() {
        return !isFirstStep();
      }
    },
    jpwFinish: {
      i18n: 'Finish &#10004;',
      show: function() {
        return isLastStep();
      }
    }
  }
};
```

## Public Methods

In general, calling methods on an element collection is preferred as it removes
any ambiguity about the walkthrough the method is affecting.

Calling methods with `$.pagewalkthrough(method, args...)` is supported, and will
operate on the currently active walkthrough if there is one. If no walkthroughs
are active, calling methods in this way will return boolean `false`.

### pagewalkthrough(options)

Creates a new walkthrough for a given element collection. Creating walkthroughs
by calling `$.pagewalkthrough` is no longer supported.

    $('body').pagewalkthrough(options);

**Note**: The walkthrough is defined on the *first* element in a collection - if
this element is removed before the walkthrough is shown, the walkthrough will
not display.

### show([name])

Starts a predefined walkthrough.

    // Show a walkthrough defined on a collection
    $('body').pagewalkthrough('show');
    // Show a walkthrough defined by a name
    $.pagewalkthrough('show', 'test');

### next()

Moves to the next step in a walkthrough, if there is one.

    // Move to the next step on a collection
    $('body').pagewalkthrough('next');
    // Move to the next step of the active walkthrough
    $.pagewalkthrough('next');

### prev()

Moves to the previous step in a walkthrough, if there is one.

    // Move to the previous step on a collection
    $('body').pagewalkthrough('prev');
    // Move to the previous step of the active walkthrough
    $.pagewalkthrough('prev')

### restart()

Moves to the first step in a walkthrough.

    // Move to the first step on a collection
    $('body').pagewalkthrough('restart');
    // Move to the first step of the active walkthrough
    $.pagewalkthrough('restart');

### close()

Closes the walkthrough.

    // Close a walkthrough on a collection
    $('body').pagewalkthrough('close');
    // Close the active walkthrough
    $.pagewalkthrough('close');

### isActive([name])

If called on the global jQuery object, the optional `name` argument
restricts the check to a specific walkthrough.

    // Returns whether *any* walkthrough is active
    $.pagewalkthrough('isActive');
    // Returns whether a specific walkthrough is active
    $.pagewalkthrough('isActive', 'test');

If called on an element collection, the `name` argument is ignored.

    // Returns whether the walkthrough defined on the collection is active
    $('body').pagewalkthrough('isActive');

### index([name])

If called on the global jQuery object, the optional `name` argument
restricts the check to a specific walkthrough.

    // Returns the current index for the active walkthrough, or false
    // if no walkthrough is active
    $.pagewalkthrough('index');
    // Returns the current index for a specific walkthrough, or false
    // if the walkthrough is not active
    $.pagewalkthrough('isActive', 'test');

If called on an element collection, the `name` argument is ignored.

    // Returns the current index, or false if the walkthrough is not active
    $('body').pagewalkthrough('index');

### getOptions([activeWalkthrough])

Returns the options for all wakthroughs, unless the `activeWalkthrough` is
`true`, in which case it returns the options for the currently active
walkthrough. If no walkthrough is active, it returns `false`.

    // Returns options for all defined walkthroughs
    $.pagewalkthrough('getOptions');
    // Returns options for the currently active walkthroughs
    $.pagewalkthrough('getOptions', true);

If called on an element collection, the `activeWalkthrough` argument is ignored
and it returns the options for the specific walkthrough.

    // Returns options for this specific walkthrough
    $('body').pagewalkthrough('getOptions');

### refresh()

Re-renders the current step, in order to handle re-positioning the overlays on
window resize or other events.

## Contributing

### Building

The `build` script in `bin/build` can be used to update the distribution files found in `dist/`.  The build script has two dependencies (see below section).
Once these two dependencies are met, just run `./bin/build` from the top-level directory of the repository to update everything in `dist/`.

#### Build script dependencies

For the build script to run, you will need the programs `less` (for compiling the LESS into CSS, and for creating the minified CSS - you'd need this
anyway if you had modified the LESS during development), `uglifyjs` (for creating the minified JS), and `jshint` (for linting the source JS).  The easiest way to install these dependencies
is through `node` and `npm`:

```shell
npm install -g less
npm install -g uglify-js
npm install -g jshint
```

### Code style

The adopted code style is that of [airbnb's JavaScript style guide](https://github.com/airbnb/javascript).  The included `.jshintrc` file is taken directly from there, with the addition of
ECMAScript 3 support (for maintaining IE9 compatibility).  Note that the **build step will fail** if `jshint` finds any errors.  You can run the linter without building by running `jshint src/`.

Also note that the CSS pre-processor [lesscss](http://lesscss.org/) is used in this project - don't modify the CSS files directly, as your changes will be overwritten when the LESS
is compiled.  Instead, you should modify the LESS and compile it (see the section on 'Building' above).

## Browser Support

### Desktop

- IE8+
- Firefox
- Safari
- Chrome (**Note: Chrome does not support cookies from locally run files.  If you want to test or develop against this aspect of the project, you should host the project on a local server**)

### Mobile

@TODO - untested as of yet

## Theme

@TODO - not yet implemented

## Changelog

#### 15/03/2016

* `v2.7.2`: #71 wasn't fixed properly - multiple resize events were causing the animation queue (and thus the callbacks) to build up - fixed by clearing queue before animating the scrolling
* `v2.7.1`: Fix bug with resize functionality - walkthrough would show when window resized even if it wasn't showing beforehand - #71; remove unsupported `lockScrolling` option from list of options - #70

#### 06/03/2016

* `v2.7.0`: Modal steps don't scroll page to top anymorei - #62; current step re-renders on resize - #68; `true` is passed to `onLeave` when walkthrough is skipped - #66

#### 20/11/2015

* `v2.6.9`: Allow next, previous and finish buttons to re-positioned using CSS, by changing which element they get appended to - #67

#### 3/11/2014

* `v2.6.8`: Add new `refresh` method, which re-renders the current step. Intended for `onresize` and similar events.
* `v2.6.7`: Clean up for code style, no API changes.
* `v2.6.6`: Added optional fallback for popup/nohighlight steps when the wrapper element cannot be found
* `v2.6.5`: The `close` method now calls `onLeave` for the current step. Return values of `false` are ignored
* `v2.6.4`: If `onEnter` or `onLeave` returns `false` for a step, move to the next/previous step as appropriate 

#### 28/08/2014

* `v2.6.3`: Bug fix for v2.6.2 which broke onBeforeShow functionality

#### 27/08/2014

* `v2.6.2`: Fix issue where plugin attempts to scroll before onBeforeShow

#### 26/08/2014

* `v2.6.1`: Fix issue where plugin wasn't scrolling to the target element properly

#### 18/08/2014

* `v2.6.0`: Add support for arrow offsets
* `v2.5.6`: Fix issue where plugin would attempt to scroll past the maximum scroll value of the container
* `v2.5.5`: Fix issue where modal steps would be mis-aligned when walkthrough re-opened
* `v2.5.4`: Fix issue where plugin would try to scroll when it shouldn't because of the `scrollTo` value being a decimal
* `v2.5.3`: Fix issue where highlight overlay could overflow the containing element

#### 14/08/2014

* `v2.5.2`: Update to test dependencies, fix centering of modal content
* `v2.5.1`: Fix position of tooltips, which broke in 2a5003f following a minor refactor
* `v2.5.0`: Remove draggable tooltip 'feature'
* `v2.4.0`: Remove defunct `accessible` and `overlay` options from step options
* `v2.3.8`: Fix clicks on tooltip content propagating through the DOM
* `v2.3.7`: Re-work of fix for issue #35, to make sure stuff behind the overlay cannot be clicked the second time a walkthrough is shown
* `v2.3.6`: Fix `onEnter` callback not firing if used with first step of a tour
* `v2.3.5`: Re-work fix for issue #36, original attempt at fixing in `v2.3.3`

#### 13/08/2014

* `v2.3.4`: Fix issue with auto-scrolling to a new target element when the element to scroll is already partly scrolled
* `v2.3.3`: Fix to prevent clicks on the overlay propagating, thus fixing issue where highlighted Bootstrap dropdowns and such would close
* `v2.3.2`: Fix overlays for popuip/tooltip content to prevent clicking things behind the walkthrough
* `v2.3.1`: Minor adjustment for more readable font-sizes
* `v2.3.0`: Fix the auto-scrolling behaviour so that it can scroll elements other than `body,html`
* `v2.2.1`: Moved the `onClose` callback to before the index reset, so close callbacks can access the last step index.

#### 12/08/2014

* `v2.2.0`: Remove support for `noHighlight` step types, add box-shadow based overlays
* `v2.1.3`: Make sure `wrapper` option selector is scoped to the current walkthrough's element, instead of being a document-wide selector
* `v2.1.2`: Fixes 2 bugs related to the `onClose` callback: 1) would not fire if walkthrough closed using `close` method, and 2) specifying an `onClose` callback would prevent
the default `close` behaviour from triggering, resulting in the walkthrough not being hidden

#### 07/08/2014

* `v2.1.1`: Fixes support for multiple walkthroughs, adds clearer method documentation and a basic test suite
* `v2.1.0`: Support for literal content in each step's `popup.content` option, instead of just a selector
* `v2.0.0`: Breaking changes to API - fix incorrect spelling of `accessable` to `accessible`; rename `stayFocus` to `lockScrolling`; remove deprecated methods
* `v1.4.0`: `name` is now a required option and **must be provided for all tours**
* `v1.3.0`: Deprecate `isPageWalkthroughActive` function in favour of `isActive` function

#### 05/08/2014

* `v1.2.4`: Add an optional finish button to the last step of the tour

#### 04/08/2014

* `v1.2.3`: Hotfix for each step's options not correctly extending default options
* `v1.2.2`: Hotfix to make the plugin actually work
* `v1.2.1`: Bug fix
* `v1.2.0`: Remove demo/example related files from master branch; source files into src/; distribution files into dist/

#### 30/07/2014

* `v1.1.4`: Add option to make close button optional
* `v1.1.3`: Support for showing next and previous buttons to move between tour stops
* `v1.1.2`: i18n support for close button text

## Contributors

### Author

* Erwin Yusrizal

### Maintainers

* James Warwood <james.duncan.1991@googlemail.com>
* Craig Roberts <craig0990@googlemail.com>

### Contributors

* Tai Nguyen
* JP Adams <jpadamsonline@gmail.com>
* James West <jwwest@gmail.com>
