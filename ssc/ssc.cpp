#include <iostream>
#include <string> 

/**
 * Sum of the digits of the paramter n
 * e.g. for 123 it returns 1+2+3 = 6
 **/
long long sum(long long n){
    long long sum = 0;
    do{
        sum+= n % 10;
        n = n / 10;
    }
    while(n>10);
    sum+= n;
    return sum;
}

/**
 * Reverse of the digits in the number n
 * e.g. for 123 it returns 321
 **/
long long inv(long long n){
    std::string s = std::to_string(n);
    long long sl = s.length();
    std::string res;
    for(n=sl-1; n>=0; n--){
        res+=s[n];
    }
    return atoll(res.c_str());
}

/**
 * be carefull... this main contains loop without a stop condition
 **/
int main(){

    long long liczba;
    long long suma, suma_b;
    for (liczba =1; ; liczba++){
        suma = sum(liczba);
        suma_b = inv(suma);
        //std::cout<<"TEST: "<< liczba<<"\t"<<suma<<"\t"<<suma_b<<std::endl;
        if(liczba == suma*suma_b){
            std::cout<<"Found number:\t"<<liczba<<" equals to: "<< suma <<" * "<< suma_b <<std::endl;
        }
        if(liczba == suma+suma_b){
            std::cout<<"Found number:\t"<<liczba<<" equals to: "<< suma <<" + "<< suma_b <<std::endl;
        }
        if(liczba == suma*suma+suma_b*suma_b){
            std::cout<<"Found number:\t"<<liczba<<" equals to: "<<suma<<"^2 + "<<suma_b<<"^2"<<std::endl;
        }
        if(liczba == suma*suma*suma_b*suma_b){
            std::cout<<"Found number:\t"<<liczba<<" equals to: "<<suma<<"^2 * "<<suma_b<<"^2"<<std::endl;
        }
    }
}
