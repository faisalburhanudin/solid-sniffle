var gulp = require('gulp');
var minifyCSS = require('gulp-minify-css');
var concatCss = require('gulp-concat-css');

var css = ['node_modules/bulma/css/bulma.css', 'assets/css/*.css'];
gulp.task('styles', function () {
    gulp.src(css)
        .pipe(minifyCSS())
        .pipe(concatCss('solid.min.css'))
        .pipe(gulp.dest('static/css'))
});

var image = ['assets/image/*'];
gulp.task('image', function () {
    gulp.src(image)
        .pipe(gulp.dest('static/image'))
});

gulp.task('watch', function () {
    gulp.watch('assets', ['styles', 'image'])
});

gulp.task('default', ['styles', 'image', 'watch']);