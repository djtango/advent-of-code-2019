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
  (reductions add origin (mapcat steps wire)))

(defn manhattan-distance [[x1 y1] [x2 y2]]
  (+ (Math/abs (- x2 x1))
     (Math/abs (- y2 y1))))

(defn -main []
  (let [[w1 w2] (parse-input "/tmp/aoc3_test")
        origin [0 0]
        w1-points (trace origin w1)
        w2-points (trace origin w2)]
    (->> [w1-points w2-points]
         (map set)
         (reduce set/intersection)
         (map (juxt identity
                    (partial manhattan-distance origin)))
         (remove #(-> % first (= [0 0])))
         (sort-by second)
         first
         second)))
