/*
Package matrix provides the Matrix type, as well as the basic operations between matrices and vectors.

There is only one Matrix type, intended to represent all sizes of matrices, as well as all vectors.

Enforcement

As of yet, there is no enforecement of dimensions when doing operations between matrices or vectors.
Until this changes, attempting interactions between mismatches matrices or vectors is considered a programming error, and its results are undefined.
In the future you can expect this to result in a Panic().
*/
package matrix
