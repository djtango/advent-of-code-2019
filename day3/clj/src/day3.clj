(ns day3
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(defn- instruction-distance [s]
  (->> s
       rest
       (apply str)
       (Integer/parseInt)))

(defn parse-input [file]
  (let [->instruction (juxt first instruction-distance)]
    (->> file
         slurp
         str/split-lines
         (map #(->> (str/split % #",")
                    (map ->instruction))))))

(def direction
  {\U [ 0  1]
   \D [ 0 -1]
   \L [-1  0]
   \R [ 1  0]})

(defn add [[x1 y1] [x2 y2]] ;; (partial mapv +) would also do
  [(+ x1 x2) (+ y1 y2)])

(defn steps [[u-d-l-r distance]]
  (repeat distance (direction u-d-l-r)))

(defn trace [origin wire]
  (->> wire
       (mapcat steps)
       (reductions add origin)
       (into [])))

(defn manhattan-distance [[x1 y1] [x2 y2]]
  (+ (Math/abs (- x2 x1))
     (Math/abs (- y2 y1))))

(defn find-best [results]
  (->> results
       (remove #(-> % first (= [0 0])))
       (sort-by second)
       first
       second))

(defn index-points [points]
  (->> points
       (map-indexed vector)
       (reduce (fn [index [steps p]]
                 (if (index p)
                   index
                   (assoc index p steps)))
               {})))

(defn get-steps-to-intersect [w1-points->steps w2-points->steps]
  (comp (partial apply +)
        (juxt w1-points->steps
              w2-points->steps)))

(defn -main []
  (let [[w1 w2] (parse-input "/tmp/aoc3")
        origin [0 0]
        w1-points (trace origin w1)
        w2-points (trace origin w2)
        w1-points->steps (index-points w1-points)
        w2-points->steps (index-points w2-points)
        intersections (->> [w1-points w2-points]
                           (map set)
                           (reduce set/intersection))]
    (->> intersections
         (map (juxt identity
                    (get-steps-to-intersect w1-points->steps
                                            w2-points->steps)))
         find-best)))
