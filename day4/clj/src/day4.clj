(ns day4)

(def bounds [138307 654504])

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
  (->> (range (first bounds)
              (inc (second bounds)))
       (filter (comp inspect str))
       (count )))
