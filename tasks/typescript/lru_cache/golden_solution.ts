/**
 * Production-grade Generic LRU Cache with TTL (Time-To-Live) support in TypeScript.
 */
export class LRUCache<K, V> {
  private capacity: number;
  private cache: Map<K, { value: V; expiresAt: number }>;

  constructor(capacity: number) {
    if (capacity <= 0) throw new Error("Capacity must be greater than zero");
    this.capacity = capacity;
    this.cache = new Map();
  }

  public get(key: K): V | undefined {
    const entry = this.cache.get(key);
    if (!entry) return undefined;

    if (Date.now() > entry.expiresAt) {
      this.cache.delete(key);
      return undefined;
    }

    // Refresh position for LRU
    this.cache.delete(key);
    this.cache.set(key, entry);
    return entry.value;
  }

  public set(key: K, value: V, ttlMs: number = 60000): void {
    if (this.cache.has(key)) {
      this.cache.delete(key);
    } else if (this.cache.size >= this.capacity) {
      // Evict oldest entry
      const oldestKey = this.cache.keys().next().value;
      if (oldestKey !== undefined) {
        this.cache.delete(oldestKey);
      }
    }

    this.cache.set(key, {
      value,
      expiresAt: Date.now() + ttlMs,
    });
  }
}
