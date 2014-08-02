module.exports = (grunt)->

  require('load-grunt-tasks')(grunt)

  grunt.registerTask('default', ['typescript', 'uglify', 'clean'])

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json')

    uglify:
      dist:
        files: 'assets/boon.min.js': ['assets/boon.js']

    typescript:
      base:
        src: ['src/ts/**/*.ts']
        dest: 'assets/boon.js'
        options:
          sourceMap: false

    watch:
      dist:
        files: ['src/ts/**/*.ts']
        tasks: ['default']
        options:
          atBegin: true

    clean: ['src/ts/**/*.js', 'assets/boon.js']

  })
