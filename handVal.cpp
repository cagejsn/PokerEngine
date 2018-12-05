#include <HandEvaluator.h>
using namespace omp;

extern "C" {
int evaluateHand(int* cards, int n){

	HandEvaluator eval;
	Hand h = Hand::empty();

	int i;
	for( i = 0 ; i < n ; i++ ) {
	
	h += Hand(cards[i]);
	
	}

	return eval.evaluate(h);


}
}




