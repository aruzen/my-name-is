import React, { useState, useCallback, useEffect, useMemo } from 'react'
import { getWords } from '../../../data/words'
import { colors, colorToHex } from '../../../data/colors'
import { WordColorAssignment, Color } from '../../../types'
import './SelectionScreen.css'

interface SelectionScreenProps {
  onComplete: (assignments: Record<string, string>) => void
  onBack: () => void
}

const SelectionScreen: React.FC<SelectionScreenProps> = ({ onComplete, onBack }) => {
  const wordsToUse = getWords()
  const [assignments, setAssignments] = useState<WordColorAssignment[]>(
    wordsToUse.map((word) => ({ word, color: null }))
  )
  const [currentWordIndex, setCurrentWordIndex] = useState(0)
  const [draggedWord, setDraggedWord] = useState<string | null>(null)
  const [circleSize, setCircleSize] = useState(() => calculateCircleSize())

  const currentWord = assignments[currentWordIndex]
  const progress = ((currentWordIndex + 1) / wordsToUse.length) * 100

  useEffect(() => {
    const handleResize = () => {
      setCircleSize(calculateCircleSize())
    }

    handleResize()
    window.addEventListener('resize', handleResize)
    return () => {
      window.removeEventListener('resize', handleResize)
    }
  }, [])

  const radius = useMemo(() => {
    const baseRadius = circleSize / 2
    const offset = Math.min(90, baseRadius * 0.35)
    return Math.max(baseRadius - offset, baseRadius * 0.55)
  }, [circleSize])

  const handleColorSelect = useCallback((color: Color) => {
    const newAssignments = [...assignments]
    newAssignments[currentWordIndex] = { ...currentWord, color }
    setAssignments(newAssignments)

    if (currentWordIndex < wordsToUse.length - 1) {
      setCurrentWordIndex(currentWordIndex + 1)
    } else {
      const result: Record<string, string> = {}
      newAssignments.forEach(({ word, color }) => {
        if (color) result[word] = color
      })
      onComplete(result)
    }
  }, [assignments, currentWordIndex, currentWord, onComplete])

  const handlePrevious = useCallback(() => {
    if (currentWordIndex > 0) {
      setCurrentWordIndex(currentWordIndex - 1)
    }
  }, [currentWordIndex])

  const handleDragStart = (e: React.DragEvent, word: string) => {
    setDraggedWord(word)
    e.dataTransfer.effectAllowed = 'move'
  }

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault()
    e.dataTransfer.dropEffect = 'move'
  }

  const handleDrop = (e: React.DragEvent, color: Color) => {
    e.preventDefault()
    if (draggedWord === currentWord.word) {
      handleColorSelect(color)
    }
    setDraggedWord(null)
  }

  return (
    <div className="selection-screen">
      <div className="selection-header">
        <button className="back-button" onClick={onBack}>
          ← 戻る
        </button>
        <div className="progress-container">
          <div className="progress-bar">
            <div 
              className="progress-fill" 
              style={{ width: `${progress}%` }}
            />
          </div>
          <span className="progress-text">
            {currentWordIndex + 1} / {wordsToUse.length}
          </span>
        </div>
      </div>

      <div className="selection-content">

        <div
          className="circular-layout"
          style={{
            width: circleSize,
            height: circleSize
          }}
        >
          <div className="word-center">
            <div
              className="current-word"
              draggable
              onDragStart={(e) => handleDragStart(e, currentWord.word)}
              onDragEnd={() => setDraggedWord(null)}
            >
              {currentWord.word}
            </div>
            {currentWord.color && (
              <div className="selected-color-info">
                <span style={{ color: colorToHex[currentWord.color] }}>
                  {currentWord.color}
                </span>
              </div>
            )}
          </div>
          
          <div className="color-circle">
            {colors.map((color, index) => {
              const angle = (index * 360) / colors.length
              const x = Math.cos((angle - 90) * Math.PI / 180) * radius
              const y = Math.sin((angle - 90) * Math.PI / 180) * radius
              
              return (
                <div
                  key={color}
                  className={`color-option color-position-${index} ${draggedWord ? 'drag-target' : ''}`}
                  style={{ 
                    backgroundColor: colorToHex[color],
                    '--x-offset': `${x}px`,
                    '--y-offset': `${y}px`
                  } as React.CSSProperties}
                  onClick={() => handleColorSelect(color)}
                  onDragOver={handleDragOver}
                  onDrop={(e) => handleDrop(e, color)}
                >
                  <span className="color-name">{color}</span>
                </div>
              )
            })}
          </div>
        </div>

        <div className="navigation-buttons">
          <button 
            className="nav-button" 
            onClick={handlePrevious}
            disabled={currentWordIndex === 0}
          >
            前の単語
          </button>
          <button 
            className="nav-button skip-button" 
            onClick={() => handleColorSelect(currentWord.color || '未選択' as Color)}
          >
            スキップ
          </button>
        </div>
      </div>
    </div>
  )
}

export default SelectionScreen

function calculateCircleSize() {
  if (typeof window === 'undefined') {
    return 500
  }

  const viewportWidth = window.innerWidth
  const viewportHeight = window.innerHeight

  const widthBased = viewportWidth - 64
  const heightBased = viewportHeight - 260
  const maxSize = 520

  const size = Math.min(maxSize, widthBased, heightBased)
  return Math.max(280, size)
}

