module.exports = function (grunt) {
    'use strict';

    // Load grunt tasks automatically
    require('load-grunt-tasks')(grunt);
    // Time how long tasks take. Can help when optimizing build times
    require('time-grunt')(grunt);

    var serveStatic = require('serve-static');

    function nocacheHeaderMiddleware(req, res, next) {
        res.setHeader('Expires', 'Thu, 01 Jan 1970 00:00:00 GMT');
        res.setHeader('Pragma', 'no-cache');
        res.setHeader('Cache-Control', 'no-store');
        next();
    }

    function getConnectConfig(folder) {
        return {
            options: {
                port: 9000,
                open: {
                    target: 'http://localhost:9000'
                },
                middleware: function (connect, options) {
                    return [nocacheHeaderMiddleware, serveStatic(folder)];
                }
            }
        };
    }

    grunt.initConfig({

        connect: {
            options: {
                port: 9000,
                hostname: 'localhost'
            },
            develop: getConnectConfig('app'),
            dist: getConnectConfig('dist')
        },

        watch: {
            options: { livereload: true },
            all: {
                files: ['app/**', '!/app/bower_components/**']
            }
        },

        karma: {
            unit: {
                configFile: 'karma.conf.js',
                singleRun: false
            },
            unit_coverage: {
                configFile: 'karma.coverage.conf.js',
                singleRun: false
            },
            unit_browser: {
                configFile: 'karma.conf.js',
                singleRun: false,
                browsers: [
                    'Chrome'
                ]
            },
            unit_ci: {
                configFile: 'karma.coverage.conf.js',
                singleRun: true
            }
        },

        protractor: {
            options: {
                configFile: 'node_modules/protractor/referenceConf.js', // Default config file
                keepAlive: true, // If false, the grunt process stops when the test fails.
                noColor: false, // If true, protractor will not use colors in its output.
                args: {
                    // Arguments passed to the command
                }
            },
            your_target: {
                options: {
                    configFile: 'protractor.conf.js', // Target-specific config file
                    args: {} // Target-specific arguments
                }
            }
        },

        wiredep: {
            index_html: {
                src: ['app/index.html']
            },
            karma_config: {
                devDependencies: true,
                src: ['karma.conf.js', 'karma.coverage.conf.js', 'Gruntfile.js'],
                ignorePath: /\.\.\//,
                fileTypes: {
                    js: {
                        block: /(([\s\t]*)\/\/\s*bower:*(\S*))(\n|\r|.)*?(\/\/\s*endbower)/gi,
                        detect: {
                            js: /'(.*\.js)'/gi
                        },
                        replace: {
                            js: '\'{{filePath}}\','
                        }
                    }
                }
            }
        },

        uglify: {
            frontend_dist_deps: {
                options: {
                    preserveComments: 'some' // preserve copyrights in minified output
                },
                files: {
                    'dist/frontend.deps.min.js': [
                        // bower:js
                        'app/bower_components/jquery/dist/jquery.js',
                        'app/bower_components/bootstrap/dist/js/bootstrap.js',
                        'app/bower_components/angular/angular.js',
                        'app/bower_components/angular-route/angular-route.js',
                        'app/bower_components/angular-ui-bootstrap-bower/ui-bootstrap-tpls.js',
                        'app/bower_components/d3/d3.js',
                        'app/bower_components/nvd3/build/nv.d3.js',
                        'app/bower_components/angular-nvd3/dist/angular-nvd3.js',
                        'app/bower_components/sprintf/src/sprintf.js',
                        // endbower
                        'tmp/templates.js'
                    ]
                }
            },
            frontend_dist: {
                options: {
                    preserveComments: 'some' // preserve copyrights in minified output
                },
                files: {
                    'dist/frontend.min.js': [
                        'app/app.js',
                        'app/app-config.js',
                        'app/shared/**/module.js',
                        'app/shared/**/*.js',
                        'app/components/**/module.js',
                        'app/components/**/*.js',
                        'tmp/templates.js'
                    ]
                }
            }

        },

        html2js: {
            options: {
                base: './app/',
                module: 'templates'
            },
            dist: {
                src: [ 'app/components/**/*.html', 'app/shared/**/*.html' ],
                dest: 'tmp/templates.js'
            }
        },

        clean: {
            tmp: {
                src: [ 'tmp' ]
            },
            dist: {
                src: [ 'dist' ]
            }
        },

        copy: {
            main: {
                files: [
                    {
                        expand: true,
                        src: '**',
                        dest: './dist/resources/images',
                        cwd: './app/resources/images'
                    },
                    {
                        expand: true,
                        src: '**',
                        dest: './dist/resources/fonts',
                        cwd: './app/bower_components/bootstrap/dist/fonts'
                    },
                    {   'dist/index.html': 'app/index-dist.html'  }
                ]
            }
        },

        cssmin: {
            options: {
                preserveComments: 'some' // preserve copyrights in minified output
            },
            target: {
                files: {
                    'dist/resources/css/frontend.min.css': [
                        'app/bower_components/bootstrap/dist/css/bootstrap.css',
                        'app/bower_components/bootstrap/dist/css/bootstrap-theme.css',
                        'app/bower_components/nvd3/build/nv.d3.css',
                        'app/resources/css/styles.css'
                    ]
                }
            }
        },

        jslint: {

            app: {
                src: [
                    'app/app.js',
                    'app/app-config.js',
                    'app/shared/**/*.js',
                    'app/components/**/*.js'
                ],
                directives: {
                    browser: true,
                    predef: [
                        'jQuery', 'angular', 'sprintf'
                    ]
                },
                options: {
                    edition: 'latest',
                    junit: 'jslint/app-junit.xml',
                    log: 'jslint/app-lint.log',
                    jslintXml: 'jslint/app-jslint.xml',
                    failOnError: false
                }
            }
        }

    });

    grunt.registerTask('test', [
        'jslint:app', 'karma:unit'
    ]);
    grunt.registerTask('test_coverage', [
        'jslint:app', 'karma:unit_coverage'
    ]);
    grunt.registerTask('test_browser', [
        'jslint:app', 'karma:unit_browser'
    ]);
    grunt.registerTask('test_ci', [
        'jslint:app', 'karma:unit_ci'
    ]);

    grunt.registerTask('serve', [
        'wiredep', 'connect:develop', 'watch'
    ]);

    grunt.registerTask('create-dist', [
        'clean:dist',
        'wiredep',
        'html2js',
        'cssmin',
        'uglify:frontend_dist_deps',
        'uglify:frontend_dist',
        'copy',
        'clean:tmp'
    ]);

    grunt.registerTask('serve-dist', [
        'create-dist',
        'connect:dist',
        'watch'
    ]);

    grunt.registerTask('serve-dist-nocreate', [
        'connect:dist',
        'watch'
    ]);

    grunt.registerTask('serve-minimal', [
        'connect:develop', 'watch'
    ]);

    grunt.loadNpmTasks('grunt-protractor-runner');
    grunt.loadNpmTasks('grunt-wiredep');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-cssmin');
    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-jslint');
    grunt.loadNpmTasks('grunt-html2js');

};
