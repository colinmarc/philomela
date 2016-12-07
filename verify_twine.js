var system = require('system'),
    page = require('webpage').create();

page.open(system.args[1], function () {
    var is_twine = page.evaluate(function() {
      return ((window['Tale'] !== undefined) ||
              (window['tale'] !== undefined) ||
              (window['story'] !== undefined)) ||
              (window['SugarCube'] !== undefined) ||
              (document.getElementsByTagName('tw-story').length > 0)
    });

    var exit_code = is_twine ? 0 : 1;
    phantom.exit(exit_code);
});
