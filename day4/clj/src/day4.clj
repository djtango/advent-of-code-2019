(ns day4
  (:require [clojure.core.async :as a :refer [go <! <!! >! >!! chan]]))

(defn bounds []
  (let [c (chan)]
    (a/onto-chan c [138307 654504])
    c))

(defn build-range [bounds]
  (let [c (chan)]
    (a/onto-chan c
                 (range (<!! bounds)
                        (inc (<!! bounds))))
    c))

(defn s<= [c1 c2]
  (if (nil? c2)
    true
    (or (neg? (compare c1 c2))
       (zero? (compare c1 c2)))))

(defn inspect [s]
  (loop [[c c' :as s] s
         [adjacent? monotonic? :as result] [false true]]
    (if (seq s)
      (recur (next s)
             [(or adjacent?
                  (= c c'))
              (and monotonic?
                   (s<= c c'))])
      (and adjacent? monotonic?))))

(defn -main []
  (let [xs (build-range (bounds))
        out (chan)]
    (<!! (a/transduce
           (filter (comp inspect str))
           (completing (fn [acc _] (inc acc)))
           0
           xs))))

(comment
  (require 'day4 :reload)
  (-main)
  )
