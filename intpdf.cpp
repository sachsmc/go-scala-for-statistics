#include <Rcpp.h>
#include <stdlib.h>
#include <math.h>
using namespace Rcpp;

// [[Rcpp::export]]
float cintpdf(int n) {

  double res = 0;
  float cx = 0;
  float cy = 0;
  float dens = 0;

  srand (time(NULL));

  double nd = (double) n;

  for(int i = 0; i < n; i++){

    cx = (float) rand()/RAND_MAX * 10 - 5;
    cy = (float) rand()/RAND_MAX * .5;

    dens = (float) exp(-cx * cx / 2.0)/sqrt(2.0 * 3.14159);

    if(cy < dens){
      res = res + 1/nd;
    }

  }

  float out =  (float) 5.0 * res;
  return out;

}
